const resultButton = document.getElementById('show-result-button');

resultButton.addEventListener('click', renderLatex)

function renderLatex(){
    const taskCode = document.getElementById('task').value;
    const answerCode = document.getElementById('answer').value;

    document.getElementById('task-box').innerHTML = '$$' + taskCode + '$$';
    document.getElementById('answer-box').innerHTML = '$$' +answerCode + '$$';

    if (window.MathJax) {
        MathJax.typesetPromise();
    }
}