// =============================
// LOGIN
// =============================

const loginForm = document.getElementById("loginForm");

if (loginForm) {

    loginForm.addEventListener("submit", async function(event) {

        event.preventDefault();

        const email = document.getElementById("email").value;
        const password = document.getElementById("password").value;

        const message = document.getElementById("message");

        message.textContent = "Logging in...";

        try {

            const response = await fetch("/login", {
                method: "POST",

                headers: {
                    "Content-Type": "application/json"
                },

                body: JSON.stringify({
                    email: email,
                    password: password
                })
            });

            const text = await response.text();

            let data;

            try {
                data = JSON.parse(text);
            } catch {
                data = {};
            }

            if (!response.ok) {

                message.textContent =
                    data.message || text || "Login failed";

                return;
            }

            // Save JWT token
            localStorage.setItem("token", data.token);

            message.textContent = "Login successful!";

            // Go to dashboard
            window.location.href = "/dashboard";

        } catch (error) {

            console.error(error);

            message.textContent =
                "Unable to connect to server.";
        }
    });
}


// =============================
// REGISTER
// =============================

const registerForm = document.getElementById("registerForm");

if (registerForm) {

    registerForm.addEventListener("submit", async function(event) {

        event.preventDefault();

        const name = document.getElementById("name").value;
        const email = document.getElementById("email").value;
        const password = document.getElementById("password").value;

        const message =
            document.getElementById("registerMessage");

        message.textContent =
            "Creating your account...";

        try {

            const response = await fetch("/register", {
                method: "POST",

                headers: {
                    "Content-Type": "application/json"
                },

                body: JSON.stringify({
                    name: name,
                    email: email,
                    password: password
                })
            });

            const text = await response.text();

            let data;

            try {
                data = JSON.parse(text);
            } catch {
                data = {};
            }

            if (!response.ok) {

                message.textContent =
                    data.message ||
                    text ||
                    "Registration failed";

                return;
            }

            message.textContent =
                "Account created successfully! Redirecting to login...";

            setTimeout(function() {

                window.location.href =
                    "/login-page";

            }, 2000);

        } catch (error) {

            console.error(error);

            message.textContent =
                "Unable to connect to server.";
        }
    });
}


// =============================
// DASHBOARD
// =============================

const balanceElement =
    document.getElementById("balance");

if (balanceElement) {

    loadAccount();
    loadTransactions();

}


// =============================
// GET ACCOUNT INFORMATION
// =============================

async function loadAccount() {

    const token =
        localStorage.getItem("token");

    const message =
        document.getElementById("dashboardMessage");

    if (!token) {

        window.location.href =
            "/login-page";

        return;
    }

    try {

        const response = await fetch(
            "/account",
            {
                method: "GET",

                headers: {
                    "Authorization":
                        "Bearer " + token
                }
            }
        );

        const text =
            await response.text();

        let data;

        try {
            data = JSON.parse(text);
        } catch {
            data = {};
        }

        if (!response.ok) {

            if (message) {
                message.textContent =
                    data.message ||
                    text ||
                    "Unable to load account.";
            }

            return;
        }


        // =============================
        // DISPLAY BALANCE
        // =============================

        const balance =
            Number(data.balance);

        document.getElementById("balance").textContent =
            new Intl.NumberFormat("en-GB", {
                style: "currency",
                currency: data.currency || "GBP",
                minimumFractionDigits: 2,
                maximumFractionDigits: 2
            }).format(balance);


        // =============================
        // DISPLAY CURRENCY
        // =============================

        document.getElementById("currency").textContent =
            data.currency || "GBP";


        // =============================
        // DISPLAY ACCOUNT NUMBER
        // =============================

        document.getElementById("accountNumber").textContent =
            data.account_number;


        // =============================
        // DISPLAY ACCOUNT TYPE
        // =============================

        document.getElementById("accountType").textContent =
            data.account_type;


    } catch (error) {

        console.error(error);

        if (message) {
            message.textContent =
                "Unable to connect to server.";
        }
    }
}


// =============================
// LOAD TRANSACTIONS
// =============================

async function loadTransactions() {

    const token =
        localStorage.getItem("token");

    const transactionsContainer =
        document.getElementById("transactions");

    if (!token) {
        return;
    }

    if (!transactionsContainer) {

        console.error(
            "Transactions container not found in dashboard.html"
        );

        return;
    }

    try {

        const response = await fetch(
            "/transactions",
            {
                method: "GET",

                headers: {
                    "Authorization":
                        "Bearer " + token
                }
            }
        );

        const text =
            await response.text();

        let data;

        try {
            data = JSON.parse(text);
        } catch {
            data = [];
        }

        if (!response.ok) {

            console.error(
                "Transactions error:",
                data
            );

            transactionsContainer.innerHTML =
                "<p>Unable to load transactions.</p>";

            return;
        }


        // =============================
        // NO TRANSACTIONS
        // =============================

        if (!Array.isArray(data) || data.length === 0) {

            transactionsContainer.innerHTML =
                "<p>No transactions yet.</p>";

            return;
        }


        // =============================
        // DISPLAY TRANSACTIONS
        // =============================

        transactionsContainer.innerHTML = "";


        data.forEach(function(transaction) {

            const transactionElement =
                document.createElement("div");

            transactionElement.className =
                "transaction";


            const amount =
                Number(transaction.amount);


            const formattedAmount =
                new Intl.NumberFormat("en-GB", {
                    style: "currency",
                    currency:
                        transaction.currency || "GBP",
                    minimumFractionDigits: 2,
                    maximumFractionDigits: 2
                }).format(amount);


            transactionElement.innerHTML = `

                <div class="transaction-info">

                    <strong>
                        ${transaction.transaction_type || "Transaction"}
                    </strong>

                    <span>
                        ${transaction.description || "Bank transaction"}
                    </span>

                    <small>
                        Reference:
                        ${transaction.reference || "N/A"}
                    </small>

                </div>

                <div class="transaction-amount">

                    <strong>
                        ${formattedAmount}
                    </strong>

                    <span>
                        ${transaction.status || "Completed"}
                    </span>

                </div>

            `;


            transactionsContainer.appendChild(
                transactionElement
            );

        });

    } catch (error) {

        console.error(
            "Unable to load transactions:",
            error
        );

        transactionsContainer.innerHTML =
            "<p>Unable to load transactions.</p>";
    }
}


// =============================
// LOGOUT
// =============================

const logoutButton =
    document.getElementById("logoutButton");

if (logoutButton) {

    logoutButton.addEventListener(
        "click",
        function() {

            // Remove JWT token
            localStorage.removeItem("token");

            // Return to login
            window.location.href =
                "/login-page";
        }
    );
}