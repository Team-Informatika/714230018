const express = require("express");
const crypto = require("crypto");
const bodyParser = require("body-parser");

const app = express();
const PORT = 3000;
const SECRET = "mysecretkey123"; // Ganti dengan Secret dari GitHub

app.use(bodyParser.json());

app.post("/webhook", (req, res) => {
    const signature = req.headers["x-hub-signature-256"];
    const payload = JSON.stringify(req.body);
    const expectedSignature = `sha256=${crypto.createHmac("sha256", SECRET).update(payload).digest("hex")}`;

    if (signature !== expectedSignature) {
        return res.status(401).json({ message: "Invalid signature" });
    }

    console.log("GitHub Webhook Received:", req.body);
    res.status(200).json({ message: "Webhook received!" });
});

app.listen(PORT, () => {
    console.log(`Webhook server running on port ${PORT}`);
});

