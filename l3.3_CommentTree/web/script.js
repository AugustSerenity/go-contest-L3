const API_BASE = "/comments";

async function loadComments() {
  const res = await fetch(`${API_BASE}/tree`);
  const data = await res.json();
  const container = document.getElementById("comments");
  container.innerHTML = "";
  renderComments(data, container);
}

function renderComments(comments, container, level = 0) {
  comments.forEach(comment => {
    const div = document.createElement("div");
    div.className = "comment";
    div.style.marginLeft = `${level * 20}px`;
    div.innerHTML = `
      <div><strong>ID ${comment.id}</strong>: ${comment.text}</div>
      <div class="actions">
        <button onclick="deleteComment(${comment.id})">Удалить</button>
        <button onclick="replyTo(${comment.id})">Ответить</button>
      </div>
    `;
    container.appendChild(div);
    if (comment.children && comment.children.length > 0) {
      renderComments(comment.children, container, level + 1);
    }
  });
}

async function addComment() {
  const text = document.getElementById("new-comment-text").value;
  const parentId = document.getElementById("parent-id").value || null;

  const res = await fetch(API_BASE, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ text, parent_id: parentId ? Number(parentId) : null }),
  });

  document.getElementById("new-comment-text").value = "";
  document.getElementById("parent-id").value = "";
  loadComments();
}

async function deleteComment(id) {
  if (!confirm("Удалить комментарий и все ответы?")) return;

  await fetch(`${API_BASE}/${id}`, { method: "DELETE" });
  loadComments();
}

function replyTo(id) {
  document.getElementById("parent-id").value = id;
  document.getElementById("new-comment-text").focus();
}

async function searchComments() {
  const q = document.getElementById("search-input").value;
  const res = await fetch(`/search?q=${encodeURIComponent(q)}`);
  const data = await res.json();
  const container = document.getElementById("comments");
  container.innerHTML = "<h3>Результаты поиска:</h3>";
  renderComments(data, container);
}
