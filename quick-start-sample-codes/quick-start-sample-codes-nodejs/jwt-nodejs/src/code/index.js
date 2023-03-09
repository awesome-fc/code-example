const express = require("express");
const jsonwebtoken = require("jsonwebtoken");

const app = express();
app.use(express.json());
app.use(express.static("static"));

app.post("/login", (req, res) => {
  const { username, password } = req.body;
  console.log(`${username} is trying to login ..`);

  // Change it to your own logic.
  if (username === "fc" && password === "fc") {
    return res.json({
      token: jsonwebtoken.sign({ user: username }, process.env.JWT_SECRET, {
        algorithm: "RS256",
        expiresIn: "1h",
      }),
    });
  }

  return res
    .status(401)
    .json({
      message:
        "The username and password your provided are invalid. Default username/password is fc/fc",
    });
});

app.get("/api/secure-resource", (req, res) => {
  res.send("Hello " + req.headers["user-id"] + " at " + new Date().toString());
});

app.listen(9000, () => {
  console.log("Example app listening on port 9000");
});
