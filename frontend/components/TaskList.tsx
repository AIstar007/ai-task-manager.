export default function TaskList({ tasks }) {
    return (
        <ul>
            {tasks.map((task) => (
                <li key={task.id}>{task.title} - {task.status}</li>
            ))}
        </ul>
    );
}
