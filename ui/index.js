const buffer = document.querySelector('.buffer');

window.onkeydown = async e => {
  if (e.ctrlKey) {
    try {
      await ctrlKeyDownBind(e.key);
    } catch (err) {
      alert(`ctrl key error: ${err}`);
    }
  } else {
    const { offset, bufferString } = await keyDownBind(e.key);
    buffer.innerHTML =
      bufferString.substr(0, offset) +
      '<div class="caret"></div>' +
      bufferString.substr(offset);
  }
};
