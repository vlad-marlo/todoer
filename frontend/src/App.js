import React from 'react';
import './App.css';
import TodoList from "./Todo/TodoList";

function App() {
    const [data, setData] = React.useState([]);
    React.useEffect(() => {
        fetch("/api/v1/tasks")
            .then((res) => res.json())
            .then((data) => setData(data.result))
            .catch((e) => {
                console.log(e)
            })
    }, []);

    return (
        <div className="wrapper">
            <h1>TODOer</h1>
            <TodoList items={data}/>
        </div>
    )
}

export default App;
