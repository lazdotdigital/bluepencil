on("keyDown", ({ key }) => {
  if (key === "{") {
    execute("insert", { text: "}" });
    execute("moveLeft");
  }
});
