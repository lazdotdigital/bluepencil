const buffer = document.querySelector('.buffer');

const renderBufferWithCaret = (offset, bufferString) => {
  buffer.innerHTML =
    bufferString.substr(0, offset) +
    '<div class="caret"></div>' +
    bufferString.substr(offset);
};

getBufferStringBind().then(s => renderBufferWithCaret(0, s));

window.onkeydown = async e => {
  e.preventDefault();
  if (e.ctrlKey) {
    try {
      await ctrlKeyDownBind(e.key);
    } catch (err) {
      alert(`ctrl key error: ${err}`);
    }
  } else {
    const { offset, bufferString } = await keyDownBind(e.key);
    renderBufferWithCaret(offset, bufferString);
  }
};
