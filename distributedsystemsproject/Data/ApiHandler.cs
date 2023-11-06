using System.Text.Json;
using System.Text.Json.Serialization;

namespace distributedsystemsproject.Data
{
    public class ApiHandler
    {
        public string URL = "http://localhost:9988";

        public async Task<string> Test()
        {
            try
            {
                using (HttpClient client = new())
                {
                    client.Timeout = new(0, 0, 10);
                    var response = await client.GetAsync(URL + "/test");
                    if (!response.IsSuccessStatusCode) return "error";
                    string json = await response.Content.ReadAsStringAsync();
                    string? todo = JsonSerializer.Deserialize<string>(json);
                    if (todo == null) return "error";
                    return todo;
                }
            }
            catch (Exception ex)
            {
                return ex.Message + " error";
            }
        }

        public async Task<TodoItem> GetElement(string id)
        {
            try
            {
                using (HttpClient client = new())
                {
                    client.Timeout = new(0, 0, 10);
                    var response = await client.GetAsync(URL + "/post/" + id);
                    if (!response.IsSuccessStatusCode) return new();
                    string json = await response.Content.ReadAsStringAsync();
                    TodoItem? todoitem = JsonSerializer.Deserialize<TodoItem>(json);
                    if (todoitem == null) return new();
                    return todoitem;
                }
            }
            catch
            {
                return new();
            }
        }

        public async Task<List<TodoItem>> GetElements()
        {
            try
            {
                using (HttpClient client = new())
                {
                    client.Timeout = new(0, 0, 10);
                    var response = await client.GetAsync(URL + "/posts");
                    if (!response.IsSuccessStatusCode) return new();
                    string json = await response.Content.ReadAsStringAsync();
                    TodoItem[]? todoarr = JsonSerializer.Deserialize<TodoItem[]>(json);
                    if (todoarr == null) return new();
                    return todoarr.ToList();
                }
            }
            catch (Exception ex)
            {
                return new();
            }
        }

        public void UpdateElement(TodoItem item)
        {
            try
            {
                using (HttpClient client = new())
                {
                    client.Timeout = new(0, 0, 10);
                }
            }
            catch { return; }
        }

        public async Task<bool> DeleteElement(string id)
        {
            try
            {
                using (HttpClient client = new())
                {
                    client.Timeout = new(0, 0, 10);
                    var response = await client.DeleteAsync(URL + "/post/" + id);
                    if (!response.IsSuccessStatusCode) return false;
                    string json = await response.Content.ReadAsStringAsync();
                    string? todoitem = JsonSerializer.Deserialize<string>(json);
                    if (todoitem == null) return false;
                    return true;
                }
            }
            catch
            {
                return false;
            }
        }

        public async Task<bool> CreateElement(string title, string desc, string color, string img)
        {
            Dictionary<string, string> keyValuePairs = new()
            {
                {"title", title},
                {"desc", desc},
                {"color", color},
                {"img", img }
            };
            FormUrlEncodedContent content = new FormUrlEncodedContent(keyValuePairs);

            try
            {
                using (HttpClient client = new())
                {
                    client.Timeout = new(0, 0, 10);
                    var response = await client.PostAsync(URL + "/post", content);
                    if (!response.IsSuccessStatusCode) return false;
                    string json = await response.Content.ReadAsStringAsync();
                    string? todoitem = JsonSerializer.Deserialize<string>(json);
                    if (todoitem == null) return false;
                    return true;
                }
            }
            catch
            {
                return new();
            }
        }
    }
}
