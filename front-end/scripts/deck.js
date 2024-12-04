let deckNames = {}; // Storage to map name of deck by ID

document.addEventListener('DOMContentLoaded', function () {
    const userID = localStorage.getItem('userID');
    if (!userID) {
        alert('User not logged in');
        window.location.href = '/index.html';
        return;
    }

    // Carregar dados das *decks* e matches ao iniciar a página
    fetch('http://localhost:8080/decks')
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            data.forEach(deck => {
                deckNames[deck.ID] = deck.DeckName; // Map deck ID to name
            });
            console.log("Deck names loaded: ", deckNames);
            loadMatches(); // Carregar a tabela de matches armazenada
        })
        .catch(error => {
            console.error('Error fetching decks:', error);
        });

    // Função para carregar a tabela de matches do localStorage
    function loadMatches() {
        const savedMatches = JSON.parse(localStorage.getItem('matches')) || [];
        savedMatches.forEach(match => {
            addRowToTable(match.userDeckInt, match.opponentDeckInt, match.wins, match.losses);
        });
    }

    function updateTable() {
        // Simulating table update with sample data
        const exampleMatches = [
            { userDeckInt: 1, opponentDeckInt: 2, wins: 5, losses: 3 },
            { userDeckInt: 2, opponentDeckInt: 1, wins: 2, losses: 6 },
        ];

        exampleMatches.forEach(match => {
            addRowToTable(match.userDeckInt, match.opponentDeckInt, match.wins, match.losses);
        });
    }

    function addRowToTable(userDeckInt, opponentDeckInt, winsInt, lossesInt) {
        const table = document.getElementById('deckStatsTable').getElementsByTagName('tbody')[0];
        const newRow = table.insertRow();

        newRow.insertCell(0).textContent = deckNames[userDeckInt] || `ID: ${userDeckInt}`;
        newRow.insertCell(1).textContent = deckNames[opponentDeckInt] || `ID: ${opponentDeckInt}`;
        newRow.insertCell(2).textContent = winsInt;
        newRow.insertCell(3).textContent = lossesInt;

        const totalMatches = winsInt + lossesInt;
        const winRate = totalMatches ? (winsInt / totalMatches) * 100 : 0; // Avoid division by zero
        newRow.insertCell(4).textContent = `${winRate.toFixed(2)}%`;

        // Salvar os dados da tabela no localStorage
        saveMatches();
    }

    // Função para salvar os dados da tabela de matches no localStorage
    function saveMatches() {
        const table = document.getElementById('deckStatsTable').getElementsByTagName('tbody')[0];
        const rows = table.getElementsByTagName('tr');
        const matches = [];

        for (let i = 0; i < rows.length; i++) {
            const row = rows[i];
            const userDeckInt = Object.keys(deckNames).find(key => deckNames[key] === row.cells[0].textContent);
            const opponentDeckInt = Object.keys(deckNames).find(key => deckNames[key] === row.cells[1].textContent);
            const wins = parseInt(row.cells[2].textContent);
            const losses = parseInt(row.cells[3].textContent);

            matches.push({
                userDeckInt,
                opponentDeckInt,
                wins,
                losses
            });
        }

        // Salvar as matches no localStorage
        localStorage.setItem('matches', JSON.stringify(matches));
    }

    // Handle deck creation
    document.getElementById('deckForm').addEventListener('submit', function (e) {
        e.preventDefault();

        const deckName = document.getElementById('deckName').value;
        const userIdInt = parseInt(userID, 10);

        if (isNaN(userIdInt)) {
            alert('Invalid user ID');
            return;
        }

        fetch('http://localhost:8080/decks', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                deck_name: deckName,
                user_id: userIdInt,
            }),
        })
            .then(response => {
                if (!response.ok) {
                    return response.text().then(text => {
                        console.log('Error response:', text);
                        throw new Error(text);
                    });
                }
                return response.json();
            })
            .then(data => {
                alert('Deck created successfully!');
                deckNames[data.deckId] = data.deck_name;
            })
            .catch(error => {
                console.error('Error:', error);
                alert('Failed to create deck: ' + error.message);
            });
    });

    // Handle WinRate updates
    document.getElementById('winrateForm').addEventListener('submit', function (e) {
        e.preventDefault();

        const userDeck = document.getElementById('userDeck').value;
        const opponentDeck = document.getElementById('opponentDeck').value;
        const wins = document.getElementById('wins').value;
        const losses = document.getElementById('losses').value;

        const userDeckInt = parseInt(userDeck);
        const opponentDeckInt = parseInt(opponentDeck);
        const winsInt = parseInt(wins);
        const lossesInt = parseInt(losses);

        fetch('http://localhost:8080/matches', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                user_deck_id: userDeckInt,
                opponent_deck_id: opponentDeckInt,
                victories: winsInt,
                defeats: lossesInt,
            }),
        })
            .then(response => {
                if (!response.ok) {
                    return response.text().then(text => {
                        console.log('Error response:', text);
                        throw new Error(text);
                    });
                }
                return response.json();
            })
            .then(data => {
                alert('WinRate updated successfully!');
                addRowToTable(userDeckInt, opponentDeckInt, winsInt, lossesInt);
            })
            .catch(error => {
                console.error('Error:', error);
            });
    });
});
