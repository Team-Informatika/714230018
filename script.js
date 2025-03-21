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

    fetch("https://t.if.co.id/714230018/webhook", { // Ganti dengan URL webhook yang valid
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
    .then(responseData => {
        console.log("Webhook Response:", responseData);
        alert("Pesan berhasil dikirim!");
    })
    .catch(error => {
        console.error("Error mengirim webhook:", error);
        alert("Gagal mengirim pesan. Silakan coba lagi nanti.");
    });
}

function getWebhookToken() {
    fetch("https://t.if.co.id/714230018/get-token") // Ganti dengan endpoint yang benar
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
