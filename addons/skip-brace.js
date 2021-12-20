on("keyDown", ({ key }) => {
  if (key === "}") {
    execute("moveLeft");
    execute("moveLeft");
    const prevChar = execute("currentChar");
    execute("moveRight");
    execute("moveRight");
    if (prevChar === "{") {
      execute("delete");
    }
  }
});
