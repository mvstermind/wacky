# **Wacky** ⏰📂 - The Goofy Go File!

Welcome to **Wacky**—the Go-powered, file-watching wizardry that keeps an eye on your project’s files and runs your specified command faster than you can say “spaghetti code”! Whether you're tweaking some code or adding random print statements (don’t lie, we all do it), Wacky Watcher’s got your back. 😎

## **Features** 🎉

- **🕵️ File Stalker:** Keeps a close watch on every file in your project—when it changes, Wacky pounces!
- **⚡ Instant Action:** Runs the command you specify as soon as it spots a change.
- **🎛️ TUI Simplicity:** No fancy GUI here—just plain ol’ terminal interface, because that’s how real developers roll.

## **Requirements** 💼

- **Go:** Because this is a Go party, and everyone’s invited.
- **Your sanity:** You might need it when you’re setting up the command to run.

## **Installation** 🚀

Clone this repository and get ready for some serious file-watching action:

```bash
git clone https://github.com/mvstermind/wacky.git
cd wacky
go build -o wacky
```

## **Usage** 🛠️

Fire up Wacky and tell it what to do when it detects a change:

```bash
./wacky -e <your-command-here>
```

For example, if you want to run tests every time you save a file:

```bash
./wacky -e go test ./...
```

Now, just sit back, relax, and let Wacky do the heavy lifting.

### **How It Works** 🧙‍♂️

1. **Watch Mode:** Wacky scans the current directory, gathering all files that aren’t trying to hide (i see you, `.git`!).
2. **On Guard:** It constantly checks for any sneaky changes. The moment something changes, it shouts “FOUND CHANGE!” in the logs.
3. **Execute!:** Your command is executed with the same precision and flair as a chef slicing veggies on a cooking show. Boom! Your command runs, and the output is displayed right there in your terminal.

## **Contributing** 🛠️

If you’ve got some cool ideas to make Wacky even wackier, or you’ve found a bug that needs squashing, feel free to fork the repo and submit a pull request. Let’s make this the best darn file watcher in Go history!

## **License** 📜

This project is licensed under the MIT License. Basically, you can do whatever you want with it, but if it breaks your project, don’t come crying to us—Wacky takes no prisoners!

---

Behold the mighty **Wacky**—your new best friend in terminal-based tomfoolery! Go forth, code warriors, and may your commands always execute flawlessly. 🚀



