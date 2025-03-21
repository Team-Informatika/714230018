function showAlert() {
    alert("Terima kasih telah menghubungi saya!");
}

function sendWebhook() {
    const data = {
        name: "Efendi Sugiantoro",
        email: "efendisugiantoro16@gmail.com",
        phone: "+6282332963807",
        message: "Seseorang menekan tombol Hubungi Saya!"
    };

    fetch("https://t.if.co.id/714230018/webhook", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }
        return response.json();
    })
    .then(data => {
        console.log("Webhook Response:", data);
        alert("Pesan berhasil dikirim!");
    })
    .catch(error => {
        console.error("Error mengirim webhook:", error);
        alert("Gagal mengirim pesan. Silakan coba lagi nanti.");
    });
}

function getWebhookToken() {
    fetch("https://.com/webhook") // Ganti dengan URL server webhook
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
