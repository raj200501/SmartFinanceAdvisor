import React, { useState, useEffect } from 'react';
import './styles.css';

function App() {
    const [financeData, setFinanceData] = useState({});
    const [userData, setUserData] = useState({});

    useEffect(() => {
        fetch('/api/finance')
            .then(response => response.json())
            .then(data => setFinanceData(data));

        fetch('/api/user')
            .then(response => response.json())
            .then(data => setUserData(data));
    }, []);

    return (
        <div className="App">
            <header className="App-header">
                <h1>Welcome, {userData.name}</h1>
                <h2>Your Financial Summary</h2>
                <p>Balance: ${financeData.balance}</p>
                <p>Expenses: ${financeData.expenses}</p>
                <p>Income: ${financeData.income}</p>
            </header>
        </div>
    );
}

export default App;
