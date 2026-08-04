const resultButton = document.getElementById('show-result-button');

resultButton.addEventListener('click', renderLatex)

function renderLatex(){
    const taskCode = document.getElementById('task').value;

    document.getElementById('task-box').innerHTML = taskCode;

    if (window.MathJax) {
        MathJax.typesetPromise();
    }
}

const url = 'http://localhost:4000/get_data';
const addTaskButton = document.getElementById('add-task')

addTaskButton.addEventListener('click', sendTask)

async function sendTask(){
    const taskCode = document.getElementById('task').value;
    const parentID = document.getElementById('chapter').value;

    const data = {
        name : taskCode,
        parent_id : Number(parentID),
        is_task : true,
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