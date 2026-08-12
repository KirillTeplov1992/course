// Константы конфигурации
const SERVER_URL = 'http://localhost:4000/get_task_with_picture';

// Элементы DOM
const resultButton = document.getElementById('show-result-button');
const addTaskButton = document.getElementById('add-task');

// Слушатели событий
resultButton.addEventListener('click', renderLatex);
addTaskButton.addEventListener('click', sendTask);

// Функция локального рендеринга
function renderLatex() {
    const taskCode = document.getElementById('task').value;
    const answerCode = document.getElementById('answer').value;
    const input = document.getElementById('imageInput');
    const image = document.getElementById('Image');

    // Безопасная обработка загрузки изображения
    if (input.files && input.files[0]) {
        const file = input.files[0];
        image.src = URL.createObjectURL(file);
        image.style.display = 'block'; // Показываем тег img, если он был скрыт
    } else {
        image.src = '';
        image.style.display = 'none'; // Скрываем, если файл не выбран
    }

    // Вставка текста формул
    document.getElementById('task-box').innerHTML = '$$' + taskCode + '$$';
    document.getElementById('answer-box').innerHTML = '$$' + answerCode + '$$';

    // Перерисовка MathJax
    if (window.MathJax && typeof MathJax.typesetPromise === 'function') {
        MathJax.typesetPromise().catch((err) => console.error('MathJax error:', err));
    }
}

async function sendTask() {
    const chapterID = document.getElementById('chapter').value;
    const taskCode = document.getElementById('task').value;
    const answerCode = document.getElementById('answer').value;
    const input = document.getElementById('imageInput');

    // Создаем объект FormData вместо обычного объекта JSON
    const formData = new FormData();
    formData.append('chapter', chapterID)
    formData.append('task', taskCode);
    formData.append('answer', answerCode);
    formData.append('type_content', 'task_with_picture');

    // Если файл выбран, добавляем его в форму
    if (input.files && input.files[0]) {
        formData.append('image', input.files[0]);
    }

    try {
        const response = await fetch(SERVER_URL, {
            method: 'POST',
            // ВАЖНО: Мы НЕ пишем headers: {'Content-Type': '...'}
            // Браузер сделает это автоматически
            body: formData 
        });

        if (!response.ok) {
            throw new Error(`Ошибка HTTP: ${response.status}`);
        }

        const result = await response.json();
        console.log('Ответ сервера:', result);
        alert('Задача и картинка успешно сохранены!');
    } catch (error) {
        console.error('Ошибка при отправке:', error);
        alert('Не удалось сохранить данные.');
    }
}