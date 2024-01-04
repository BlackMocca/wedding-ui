window.execCommandCopy = (textToCopy) => {
    const textArea = document.createElement("textarea");
    textArea.value = textToCopy;
    document.body.appendChild(textArea);
    
    textArea.select();
    document.execCommand("copy");
    
    document.body.removeChild(textArea);
}