function openConsole() {
    document.getElementById('console-modal').style.display = 'flex';
    showLoginForm();
}

function closeConsole() {
    document.getElementById('console-modal').style.display = 'none';
    document.getElementById('console-content').innerHTML = '';
}

function showLoginForm() {
    document.getElementById('console-content').innerHTML = `
        <h2>管理员登录</h2>
        <form id="login-form" onsubmit="return loginAdmin(event)">
            <input type="text" id="username" placeholder="用户名" value="admin" required style="width:90%;margin-bottom:10px;padding:8px;">
            <input type="password" id="password" placeholder="密码" value="admin" required style="width:90%;margin-bottom:10px;padding:8px;">
            <button type="submit" style="width:100%;padding:10px;background:#1976d2;color:#fff;border:none;border-radius:4px;">登录</button>
        </form>
        <div id="login-error" style="color:red;margin-top:8px;"></div>
    `;
}

function loginAdmin(e) {
    e.preventDefault();
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    fetch('/api/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password })
    })
    .then(res => res.json().then(data => ({status: res.status, data})))
    .then(({status, data}) => {
        if(status === 200) {
            loadConfigPanel();
        } else {
            document.getElementById('login-error').innerText = data.error || '登录失败';
        }
    });
    return false;
}

function loadConfigPanel() {
    fetch('/api/admin/config', { credentials: 'include' })
    .then(res => {
        if(res.status === 401) {
            showLoginForm();
            return;
        }
        return res.json();
    })
    .then(cfg => {
        if(!cfg) return;
        document.getElementById('console-content').innerHTML = `
            <h2>服务配置</h2>
            <pre id="config-json" style="background:#f5f5f5;padding:12px;border-radius:6px;">${JSON.stringify(cfg, null, 2)}</pre>
            <button onclick="showEditConfig()" style="margin-right:10px;">编辑配置</button>
            <button onclick="resetConfig()" style="background:#e53935;color:#fff;">恢复默认</button>
        `;
    });
}

function showEditConfig() {
    fetch('/api/admin/config', { credentials: 'include' })
    .then(res => res.json())
    .then(cfg => {
        document.getElementById('console-content').innerHTML = `
            <h2>编辑配置</h2>
            <textarea id="edit-config" style="width:100%;height:200px;">${JSON.stringify(cfg, null, 2)}</textarea>
            <div style="margin-top:10px;">
                <button onclick="saveConfig()" style="margin-right:10px;">保存</button>
                <button onclick="loadConfigPanel()">取消</button>
            </div>
            <div id="save-error" style="color:red;margin-top:8px;"></div>
        `;
    });
}

function saveConfig() {
    let val = document.getElementById('edit-config').value;
    let cfg;
    try {
        cfg = JSON.parse(val);
    } catch(e) {
        document.getElementById('save-error').innerText = 'JSON 格式错误';
        return;
    }
    fetch('/api/admin/config', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify(cfg)
    })
    .then(res => res.json().then(data => ({status: res.status, data})))
    .then(({status, data}) => {
        if(status === 200) {
            loadConfigPanel();
        } else {
            document.getElementById('save-error').innerText = data.error || '保存失败';
        }
    });
}

function resetConfig() {
    if(!confirm('确定要恢复为系统默认配置吗？')) return;
    fetch('/api/admin/config/reset', {
        method: 'POST',
        credentials: 'include'
    })
    .then(res => res.json())
    .then(data => {
        if(data.msg) {
            loadConfigPanel();
        } else {
            alert(data.error || '恢复失败');
        }
    });
}
