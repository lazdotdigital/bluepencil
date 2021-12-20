const buffer = document.querySelector('.buffer');

const renderBuffer = (offset, bufferString) => {
  const withCaret =
    bufferString.substr(0, offset) +
    '<div class="caret"></div>' +
    bufferString.substr(offset);

  buffer.innerHTML = withCaret
    .split('\n')
    .map((line, i) => `<span class="line-number">${i}</span>${line}`)
    .join('\n');
};

getBufferStringBind().then(s => renderBuffer(0, s));

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
    renderBuffer(offset, bufferString);
  }
};
