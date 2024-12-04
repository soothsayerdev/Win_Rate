// Função para lidar com o registro de usuário
document.getElementById('registerForm').addEventListener('submit', function (e) {
    e.preventDefault();

    const email = document.getElementById('registerEmail').value;
    const password = document.getElementById('registerPassword').value;

    const user = {
        email: email,
        password: password
    };

    // Enviar os dados para o backend para registrar o usuário
    fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(user),
    })
    .then(response => response.json())
    .then(data => {
        if (data.message) {
            alert(data.message); // Exibe a mensagem de sucesso
            // Redireciona para a tela de login após o registro
            window.location.href = '/login.html';
        }
    })
    .catch(error => {
        console.error('Erro ao registrar usuário:', error);
        alert('Erro ao registrar usuário. Tente novamente.');
    });
});

// Função para lidar com o login do usuário
document.getElementById('loginForm').addEventListener('submit', function (e) {
    e.preventDefault();

    const email = document.getElementById('loginEmail').value;
    const password = document.getElementById('loginPassword').value;

    const user = {
        email: email,
        password: password
    };

    // Enviar os dados para o backend para autenticar o usuário
    fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(user),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Credenciais inválidas');
        }
        return response.json();
    })
    .then(data => {
        if (data.message) {
            alert(data.message); // Exibe a mensagem de sucesso
            // Armazenar o token no localStorage ou o que for retornado para autenticação
            localStorage.setItem('userEmail', email);
            window.location.href = '/home.html'; // Redireciona para a página inicial após login
        }
    })
    .catch(error => {
        console.error('Erro ao fazer login:', error);
        alert('Erro ao fazer login. Verifique suas credenciais e tente novamente.');
    });
});
