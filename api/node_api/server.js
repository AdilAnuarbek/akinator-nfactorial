const express = require('express');
const cors = require('cors');
const { AkinatorClient, Languages, Themes, Answers } = require("node_akinator");

const app = express();
app.use(cors());
app.use(express.json());

const sessions = new Map();
const regions = {
    "en": Languages.English,
    "ar": Languages.Arabic,
    "he": Languages.Hebrew,
    "jp": Languages.Japanese,
    "kr": Languages.Korean,
    "nl": Languages.Dutch,
    "pl": Languages.Polish,
    "id": Languages.Indonesian,
    "fr": Languages.French,
    "es": Languages.Spanish,
    "de": Languages.German,
    "it": Languages.Italian,
    "pt": Languages.Portuguese,
    "ru": Languages.Russian,
    "tr": Languages.Turkish,
    "cn": Languages.Chinese,
};

const themes = {
    "characters": Themes.Character,
    "animals": Themes.Animals,
    "objects": Themes.Objects,
};

app.post('/start', async (req, res) => {
    try {
        const { sessionId, theme, region } = req.body;

        const Theme = themes[theme];
        const Region = regions[region];

        if (!sessionId) return res.status(400).json({ error: 'sessionId required' });

        const aki = new AkinatorClient(Region, true, Theme);
        const start = await aki.start();  // begin game
        sessions.set(sessionId, aki);
        res.json({
            question: start.question,  // first question
        });
    } catch (err) {
        console.error(err);
        res.status(500).json({ error: 'failed to start akinator' });
    }
});

const answers = {
    "1": Answers.No,
    "0": Answers.Yes,
    "2": Answers.IDontKnow,
    "3": Answers.Probably,
    "4": Answers.ProbablyNot,
}

/**
 * POST /answer
 * Body: { sessionId: string, answer: number }
 * answer is the index in the 'answers' array returned previously
 */
app.post('/answer', async (req, res) => {
    try {
        const { sessionId, answer } = req.body;
        const Answer = answers[answer];
        const aki = sessions.get(sessionId);
        if (!aki) return res.status(404).json({ error: 'session not found' });

        // AnswerResult
        const answer_returned = await aki.answer(Answer);
        if (answer_returned.won) {
            res.json({ 
                question: "", 
                win: true, 
                name: aki.winResult.name 
            });
            return;
        }

        // Otherwise, send next question
        res.json({
            question: answer_returned.question,
            win: false,
            name: "",
        });
    } catch (err) {
        console.error(err);
        res.status(500).json({ error: 'failed to process answer' });
    }
});

app.post('/guess', async (req, res) => {
    try {
        const { sessionId, answer } = req.body;
        const aki = sessions.get(sessionId);
        if (!aki) return res.status(404).json({ error: 'session not found' });
        // AnswerResult
        if (answer === "1") {
            await aki.submitWin();
            sessions.delete(sessionId);
            res.json({
                question: "",
            })
            return;
        } if (answer === "0") {
            const answer = await aki.continue();
            res.json({
                question: answer.question,
            });
        }
    } catch (err) {
        console.error(err);
        res.status(500).json({ error: 'failed to process answer' });
    }
})

app.listen(3000, () => console.log("Server ready on port 3000."));