function showAlert() {
    alert("Terima kasih telah menghubungi saya!");
    sendWebhook();
}

function sendWebhook() {
    const data = {
        name: "Efendi Sugiantoro",
        email: "efendisugiantoro16@gmail.com",
        phone: "+6282332963807",
        message: "Seseorang menekan tombol Hubungi Saya!"
    };

    fetch("https://asia-southeast2-awangga.cloudfunctions.net/domyid/webhook/github/pemograman_3_2a", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "X-Hub-Signature": "sha256=" + btoa("T4IAwwA5gv8MJyp6GSCtaTocMyTmUwpBF3nVuVUzVysIpWUx")
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }
        return response.json();
    })
    .then(responseData => {
        console.log("Webhook Response:", responseData);
        alert("Pesan berhasil dikirim ke webhook!");
    })
    .catch(error => {
        console.error("Error mengirim webhook:", error);
        alert("Gagal mengirim pesan. Silakan coba lagi nanti.");
    });
}

function getWebhookToken() {
    fetch("https://asia-southeast2-awangga.cloudfunctions.net/domyid/webhook/github/get-token")
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            document.getElementById("webhook-token").innerText = data.token || "Token tidak ditemukan";
        })
        .catch(error => {
            console.error("Error mengambil token webhook:", error);
            document.getElementById("webhook-token").innerText = "Gagal mengambil token";
        });
}

// Panggil saat halaman dimuat
document.addEventListener("DOMContentLoaded", getWebhookToken);
