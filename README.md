# Blue Pencil

A simple, fast, and extendable code editor. Built in Go.

_NOTE: THIS PROJECT IS IN ITS VERY EARLY STAGES. NOT ALL FEATURES ARE COMPLETE, AND EVERYTHING IS SUBJECT TO CHANGE._

- **Snappy** text editing.
- Easy **addon** creation.
- **Simple** and **configurable** interface.

## Tutorial

### Installation

1. Clone: `git clone https://github.com/lazdotdigital/bluepencil.git`.
2. Build: `go build .`.
3. Run: `./bluepencil <file>`.

### Navigation

Edit your file as desired, then `ctrl+s` to save. The caret is moved via the arrow keys.

# Configuration

Blue Pencil is configured via environment variables. When run, the `.env` file, if present, will be loaded. The following configuration options are currently avaliable:

1. `ADDONS_PATH`

## Addons

Blue Pencil has a simple and powerful addon system. All addons are written in ES6 compliant JavaScript.

### Installation

Installing addons is very simple. All files located at `ADDONS_PATH` loaded and executed. To install one, simply drop it into the directory.

### Default Addons

There are two default addons, usable by setting the `ADDONS_PATH` environment variable to `./addons`. I suggest putting the rest of your addons in this directory, too. Feel free to move it if necessary. They are as follows:

1. `close-braces.js` automatically closes curly braces (`{}`), returning your cursor to the middle of them.
2. `skip-brace.js` automatically skips over `}` if `{` is to the left of it.

### Writing Addons

Currently, addons have the following three functions exposed to them: `log`, `on`, and `execute`.

#### log

`log` is very simple. It takes a single `string` argument and outputs it to stdout.

#### on

`on` is a more complex function that is responsible for registering event handlers. It takes two arguments: a `string` containing the event's name, and a `function` that takes an `object` describing the event and optionally throws an error.

It currently exposes the following events:

1. `'keyDown'` is fired _after_ the editor's key down handler fires. The event has a single field, `key` containing a textual representation of the key pressed.

#### execute

`execute` sends commands to the editor. It takes two arguments: a `string` containing the command's name and, if needed, an `object` describing the command's arguments. It may return values.

The following commands are avaliable:

1. `'moveUp'` moves the caret up.
2. `'moveDown'` moves the caret up.
3. `'moveLeft'` moves the caret up.
4. `'moveRight'` moves the caret up.
5. `'delete'` deletes the current character.
6. `'currentChar'` returns the current character.
7. `'insert'` inserts text into the buffer. It has one field, `text`, containing the text to insert.

### Example Addon

The following is a very simple addon that automatically closes curly braces:

```javascript
on("keyDown", ({ key }) => {
  if (key === "{") {
    execute("insert", { text: "}" });
    execute("moveLeft");
  }
});
```

## Final Notes

Thank you for taking the time to check out my editor. Please open up an issue for any concerns or questions you may have!
