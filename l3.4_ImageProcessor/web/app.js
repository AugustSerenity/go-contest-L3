const form = document.getElementById('form');
const list = document.getElementById('list');
const items = new Map();

form.addEventListener('submit', async (e) => {
    e.preventDefault();
    const fd = new FormData(form);
    const res = await fetch('/upload', { method: 'POST', body: fd });
    const json = await res.json();
    addItem(json.id);
});

function addItem(id) {
    items.set(id, { status: 'pending' });
    render();
    poll(id);
}

async function poll(id) {
    const t = setInterval(async () => {
        const r = await fetch(`/status/${id}`);
        if (!r.ok) return;
        const j = await r.json();
        items.set(id, j);
        render();
        if (j.status === 'completed') clearInterval(t);
    }, 1500);
}

function render() {
    list.innerHTML = '';
    for (const [id, it] of items) {
        const el = document.createElement('div');
        el.className = 'card';
        el.innerHTML = `
            <div class="muted">#${id}</div>
            ${it.status === 'completed' ? `<img src="/image/${id}?variant=thumb"/>` : `<div class="muted">в обработке…</div>`}
            <div>
                <a href="/image/${id}?variant=resized" target="_blank">Скачать</a>
                <a href="/image/${id}?variant=watermarked" target="_blank" style="margin-left:8px">Водяной знак</a>
            </div>
            <button data-id="${id}">Удалить</button>
        `;
        el.querySelector('button').onclick = async () => {
            await fetch(`/image/${id}`, { method: 'DELETE' });
            items.delete(id);
            render();
        };
        list.appendChild(el);
    }
}