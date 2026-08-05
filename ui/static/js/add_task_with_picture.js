const resultButton = document.getElementById('show-result-button');

resultButton.addEventListener('click', renderLatex)

function renderLatex(){
    const taskCode = document.getElementById('task').value;
    const answerCode = document.getElementById('answer').value;
    const input = document.getElementById('imageInput');
    const image = document.getElementById('Image');

    const file = input.files[0];
    const imageURL = URL.createObjectURL(file);

    document.getElementById('task-box').innerHTML = '$$' + taskCode + '$$';
    image.src = imageURL;
    document.getElementById('answer-box').innerHTML = '$$' +answerCode + '$$';

    if (window.MathJax) {
        MathJax.typesetPromise();
    }
}

const url = 'http://localhost:4000/get_data';
const addTaskButton = document.getElementById('add-task')

addTaskButton.addEventListener('click', sendTask)

async function sendTask(){
    const taskCode = document.getElementById('task').value;
    const answerCode = document.getElementById('answer').value;
    const parentID = document.getElementById('chapter').value;

    const data = {
        name : taskCode,
        answer : answerCode,
        parent_id : Number(parentID),
        type_content : 'task',
    };

    try {
    // Отправляем POST запрос на наш Go-сервер
    const response = await fetch(url, {
        method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
    });

    const result = await response.json();
    console.log('Ответ сервера:', result);
        } catch (error) {
            console.error('Ошибка:', error);
    }
}