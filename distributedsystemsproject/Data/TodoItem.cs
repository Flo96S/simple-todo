namespace distributedsystemsproject.Data
{
    public class TodoItem
    {
        public string Id { get; set; }
        public string Title { get; set; }
        public string Description { get; set; }
        public DateTime Created { get; set; }
        public bool IsDone { get; set; }
        public string Image {  get; set; }
        public string Color { get; set; }
    }

    public class TodoItemShort
    {
        public string Id { get; set; }
        public string Title { get; set; }
        public string Color { get; set; }
        public bool IsDone { get; set; }
    }
}
