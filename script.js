// Fungsi menampilkan alert dan mengirim webhook
async function showAlert() {
    alert("Terima kasih telah menghubungi saya!");
    await sendWebhook();
}

// Fungsi mengirim data ke webhook
async function sendWebhook() {
    const data = {
        name: "Efendi Sugiantoro",
        email: "efendisugiantoro16@gmail.com",
        phone: "+6282332963807",
        message: "Seseorang menekan tombol Hubungi Saya!"
    };

    try {
        const response = await fetch("https://asia-southeast2-awangga.cloudfunctions.net/domyid/webhook/github/pemograman_3_2a", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "X-Hub-Signature": "sha256=" + btoa("T4IAwwA5gv8MJyp6GSCtaTocMyTmUwpBF3nVuVUzVysIpWUx")
            },
            body: JSON.stringify(data)
        });

        if (!response.ok) throw new Error(`HTTP error! Status: ${response.status}`);
        
        const result = await response.json();
        console.log("Webhook Response:", result);
        alert("Pesan berhasil dikirim ke webhook!");
    } catch (error) {
        console.error("Error mengirim webhook:", error);
        alert("Gagal mengirim pesan. Silakan coba lagi nanti.");
    }
}

// Fungsi mengambil token webhook
async function getWebhookToken() {
    const tokenElement = document.getElementById("webhook-token");
    tokenElement.innerText = "Mengambil token...";

    try {
        const response = await fetch("https://asia-southeast2-awangga.cloudfunctions.net/domyid/webhook/github/get-token");
        if (!response.ok) throw new Error(`HTTP error! Status: ${response.status}`);

        const data = await response.json();
        tokenElement.innerText = data.token || "Token tidak ditemukan";
    } catch (error) {
        console.error("Error mengambil token webhook:", error);
        tokenElement.innerText = "Gagal mengambil token";
    }
}

// Fungsi mengecek apakah menu.html ada sebelum diarahkan
async function checkMenuPage(event) {
    event.preventDefault(); // Mencegah navigasi sebelum pengecekan
    const menuUrl = "menu.html";

    try {
        const response = await fetch(menuUrl, { method: "HEAD" });
        if (response.ok) {
            window.location.href = menuUrl;
        } else {
            window.location.href = "404.html"; // Redirect ke halaman 404 jika tidak ditemukan
        }
    } catch (error) {
        window.location.href = "404.html"; // Redirect jika terjadi error
    }
}

// Event listener untuk tombol "Hubungi Saya"
document.getElementById("contact-btn").addEventListener("click", showAlert);

// Event listener untuk tombol "Menu"
document.getElementById("menu-btn").addEventListener("click", checkMenuPage);

// Panggil saat halaman dimuat
document.addEventListener("DOMContentLoaded", getWebhookToken);
