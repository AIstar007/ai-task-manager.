import { useEffect, useState } from "react";

export default function Dashboard() {
    const [tasks, setTasks] = useState([]);

    useEffect(() => {
        fetch("/api/tasks")
            .then((res) => res.json())
            .then(setTasks);
    }, []);

    return (
        <div className="p-4">
            <h1 className="text-xl font-bold">Task Dashboard</h1>
            <ul>
                {tasks.map((task) => (
                    <li key={task.id}>{task.title} - {task.status}</li>
                ))}
            </ul>
        </div>
    );
}
