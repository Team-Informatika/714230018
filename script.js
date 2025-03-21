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
    .then(response => response.json())
    .then(data => {
        console.log("Webhook Response:", data);
        alert("Pesan berhasil dikirim!");
    })
    .catch(error => {
        console.error("Error:", error);
        alert("Gagal mengirim pesan.");
    });
}
