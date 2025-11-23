document.addEventListener('DOMContentLoaded', () => {
    fetch('/api/packages')
        .then(response => response.json())
        .then(packages => {
            const packageList = document.getElementById('package-list').getElementsByTagName('tbody')[0];
            packages.forEach(pkg => {
                const row = packageList.insertRow();
                row.innerHTML = `
                    <td>${pkg.name}</td>
                    <td>${pkg.version}</td>
                    <td>${pkg.description}</td>
                    <td><button>Install</button></td>
                `;
            });
        });
});
