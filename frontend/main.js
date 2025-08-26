var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g = Object.create((typeof Iterator === "function" ? Iterator : Object).prototype);
    return g.next = verb(0), g["throw"] = verb(1), g["return"] = verb(2), typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var _this = this;
var API_URL = "http://localhost:8080";
// Elements
var usernameInput = document.getElementById("username");
var emailInput = document.getElementById("email");
var createUserBtn = document.getElementById("createUserBtn");
var titleInput = document.getElementById("title");
var authorInput = document.getElementById("author");
var userIdInput = document.getElementById("userId");
var createBookBtn = document.getElementById("createBookBtn");
var getUsersBtn = document.getElementById("getUsersBtn");
var getBooksBtn = document.getElementById("getBooksBtn");
var usersList = document.getElementById("usersList");
var booksList = document.getElementById("booksList");
// Create User
createUserBtn.addEventListener("click", function () { return __awaiter(_this, void 0, void 0, function () {
    var name, email, res, data;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                name = usernameInput.value;
                email = emailInput.value;
                if (!name || !email)
                    return [2 /*return*/, alert("Enter name & email")];
                return [4 /*yield*/, fetch("".concat(API_URL, "/users"), {
                        method: "POST",
                        headers: { "Content-Type": "application/json" },
                        body: JSON.stringify({ name: name, email: email })
                    })];
            case 1:
                res = _a.sent();
                return [4 /*yield*/, res.json()];
            case 2:
                data = _a.sent();
                alert("User created: ".concat(data.name, " (ID: ").concat(data.id, ")"));
                return [2 /*return*/];
        }
    });
}); });
// Create Book
createBookBtn.addEventListener("click", function () { return __awaiter(_this, void 0, void 0, function () {
    var title, author, userId, res, data;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                title = titleInput.value;
                author = authorInput.value;
                userId = parseInt(userIdInput.value);
                if (!title || !author || isNaN(userId))
                    return [2 /*return*/, alert("Fill all book fields")];
                return [4 /*yield*/, fetch("".concat(API_URL, "/books"), {
                        method: "POST",
                        headers: { "Content-Type": "application/json" },
                        body: JSON.stringify({ title: title, author: author, user_id: userId })
                    })];
            case 1:
                res = _a.sent();
                return [4 /*yield*/, res.json()];
            case 2:
                data = _a.sent();
                alert("Book created: ".concat(data.title, " for user ").concat(data.user_id));
                return [2 /*return*/];
        }
    });
}); });
// Get Users
getUsersBtn.addEventListener("click", function () { return __awaiter(_this, void 0, void 0, function () {
    var res, users;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: return [4 /*yield*/, fetch("".concat(API_URL, "/users"))];
            case 1:
                res = _a.sent();
                return [4 /*yield*/, res.json()];
            case 2:
                users = _a.sent();
                usersList.innerHTML = "";
                users.forEach(function (u) {
                    var li = document.createElement("li");
                    li.className = "border rounded p-2";
                    li.textContent = "ID: ".concat(u.id, ", Name: ").concat(u.name, ", Email: ").concat(u.email);
                    usersList.appendChild(li);
                });
                return [2 /*return*/];
        }
    });
}); });
// Get Books
getBooksBtn.addEventListener("click", function () { return __awaiter(_this, void 0, void 0, function () {
    var res, books;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: return [4 /*yield*/, fetch("".concat(API_URL, "/books"))];
            case 1:
                res = _a.sent();
                return [4 /*yield*/, res.json()];
            case 2:
                books = _a.sent();
                booksList.innerHTML = "";
                books.forEach(function (b) {
                    var li = document.createElement("li");
                    li.className = "border rounded p-2";
                    li.textContent = "ID: ".concat(b.id, ", Title: ").concat(b.title, ", Author: ").concat(b.author, ", UserID: ").concat(b.user_id);
                    booksList.appendChild(li);
                });
                return [2 /*return*/];
        }
    });
}); });
