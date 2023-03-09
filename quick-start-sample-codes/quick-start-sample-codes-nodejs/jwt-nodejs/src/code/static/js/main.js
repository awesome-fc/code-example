var interval;

function fetchData() {
  fetch("/api/secure-resource", {
    headers: {
      Authentication: "Bearer " + sessionStorage.getItem("TOKEN"),
    },
  }).then((res) => {
    var message = document.getElementById("message");
    if (res.status === 200) {
      res.text().then((text) => {
        message.innerHTML = text;
        showLoginView(false);
      });
    } else {
      showLoginView(true);
    }
  });
}

function showLoginView(show) {
  var login = document.getElementById("login");
  login.style.display = show ? "flex" : "none";
  var success = document.getElementById("success");
  success.style.display = show ? "none" : "block";
  var message = document.getElementById("message");

  if (show) {
    clearInterval(interval);
    sessionStorage.removeItem("TOKEN");
    message.innerHTML = "";
  }
}

function login() {
  var username = document.querySelector("#username").value;
  var password = document.querySelector("#password").value;
  var loginBtn = document.querySelector("#login-btn");

  if (username && password) {
    loginBtn.innerHTML = "Logging in...";
    fetch("/login", {
      method: "POST",
      mode: "cors",
      cache: "no-cache",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      redirect: "follow",
      body: JSON.stringify({
        username,
        password,
      }),
    })
      .then((res) => {
        if (res.status === 200) {
          res.json().then((json) => {
            const { token } = json;
            if (token) {
              sessionStorage.setItem("TOKEN", token);
              showLoginView(false);
              fetchData();
            }
          });
        } else {
          loginBtn.innerHTML = "Login";
          setTimeout(() => {
            alert("Username or Password is invalid.");
          }, 100);
        }
      })
      .finally(() => {
        loginBtn.innerHTML = "Login";
      });
  }
}

function logout() {
  showLoginView(true);
}

function start() {
  clearInterval(interval);
  interval = setInterval(fetchData, 3000);
  fetchData();
}

start();
