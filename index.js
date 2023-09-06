const express = require("express")

const app = express()

app.get("/", (req, res) => {
    res.json({msg: "Hi there"})
})

app.listen(3333, () => console.log("your express.js app is running."))
