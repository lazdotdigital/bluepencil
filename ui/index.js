const buffer = document.querySelector('.buffer');

window.onkeydown = e => {
  keyDownBind(e.key).then(({ offset, bufferString }) => {
    buffer.innerHTML =
      bufferString.substr(0, offset) +
      '<div class="caret"></div>' +
      bufferString.substr(offset);
  });
};
