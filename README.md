# Uy vazifasi: Go-da HTTPS-dan foydalanib RESTful API yaratish

## Maqsad
`Go`-da `HTTPS`-dan foydalanib, xavfsiz muloqot qilishni ta'minlaydigan oddiy `RESTful API` yaratish. `API` oddiy `CRUD` amallarini boshqarishi va `HTTPS` orqali ishlashi kerak.

## Talablar
1. **Sertifikatlarni yaratish**: Local ishlatishga uchun `OpenSSL` yoki shunga o'xshash vositalardan foydalangan holda o'zingizga o'zingiz imzolagan `SSL/TLS` sertifikatlarini yarating.

2. **HTTPS Server**: Set up an `HTTPS` server in `Go`, utilizing the `net/http` package to listen on a secure port.
    
3. **RESTful Endpoints:**
    - Kamida to'rtta `CRUD` endpointlarni (masalan, `GET /items`, `POST /items`, `PUT /items/:id`, `DELETE /items/:id`) amalga oshiring.
    - Noto'g'ri so'rovlar yoki resurslarning yo'qligi uchun tegishli xatolarni boshqarishni ta'minlang.

4. **Ma'lumot saqlashini:**
    - Ma'lumotlarni oddiy xotira tarkibi (masalan, map) yoki tanlagan ma'lumotlar bazasida (masalan, PostgreSQL, MongoDB va boshqalar) saqlang.

5. **Xavfsiz So'rovlar:**
    - `API`-ga barcha so'rovlar `HTTPS` orqali amalga oshirilishini ta'minlang va xavfsiz bo'lmagan `HTTP` so'rovlarini rad eting.

6. **Environment variables:**
    - Port raqami, sertifikat yo'llari va boshqa sezgir ma'lumotlarni `environment variable` orqali sozlang.

7. **Graceful Shutdown**:
    - `Go`-ning `context` paketidan foydalanib, serverni `graceful shutdown` bilan ta'minlang.