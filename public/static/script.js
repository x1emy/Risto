document.getElementById('registerForm').addEventListener('submit', async function (e) {
    e.preventDefault();

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    const message = document.getElementById('responseMessage');
    const submitButton = document.querySelector('button[type="submit"]');

    // Проверка на пустые поля
    if (!username || !password) {
        message.textContent = '❌ Все поля обязательны для заполнения!';
        return;
    }

    // Блокируем кнопку отправки
    submitButton.disabled = true;
    message.textContent = 'Загрузка...';  // Показать пользователю, что идет запрос

    try {
        const response = await fetch('/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
        });

        const result = await response.json();

        if (response.ok) {
            message.textContent = '✅ Успешная регистрация!';
            window.location.href = '/';  // Переход на главную страницу после регистрации
        } else {
            message.textContent = `❌ Ошибка: ${result.error || response.statusText}`;
        }
    } catch (error) {
        console.error('Ошибка при отправке запроса:', error);
        message.textContent = '❌ Ошибка регистрации!';
    } finally {
        // Включаем кнопку отправки обратно
        submitButton.disabled = false;
    }
});
