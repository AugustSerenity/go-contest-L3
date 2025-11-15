const EventAPI = {
    baseUrl: '',

    async fetchEvents() {
        try {
            const res = await fetch(`${this.baseUrl}/events`);
            const data = await res.json();
            if (!res.ok) throw data;
            return data;
        } catch (err) {
            console.error('Ошибка fetchEvents:', err);
            return { error: err.error || 'Ошибка сети' };
        }
    },

    async bookEvent(eventID, seats) {
        try {
            const res = await fetch(`${this.baseUrl}/events/${eventID}/book`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ seats })
            });
            const data = await res.json();
            if (!res.ok) throw data;
            return data;
        } catch (err) {
            console.error('Ошибка bookEvent:', err);
            return { error: err.error || 'Ошибка сети' };
        }
    },

    async createEvent({ name, date, capacity, paymentTTL }) {
        try {
            const res = await fetch(`${this.baseUrl}/events`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name, date, capacity, paymentTTL })
            });
            const data = await res.json();
            if (!res.ok) throw data;
            return data;
        } catch (err) {
            console.error('Ошибка createEvent:', err);
            return { error: err.error || 'Ошибка сети' };
        }
    },

    async fetchBookings(eventID) {
        try {
            const res = await fetch(`${this.baseUrl}/events/${eventID}/bookings`);
            const data = await res.json();
            if (!res.ok) throw data;
            return data;
        } catch (err) {
            console.error('Ошибка fetchBookings:', err);
            return { error: err.error || 'Ошибка сети' };
        }
    },

    async confirmBooking(bookingID) {
        try {
            const res = await fetch(`${this.baseUrl}/events/${bookingID}/confirm`, { 
                method: 'POST' 
            });
            const data = await res.json();
            if (!res.ok) throw data;
            return data;
        } catch (err) {
            console.error('Ошибка confirmBooking:', err);
            return { error: err.error || 'Ошибка сети' };
        }
    },

    renderEvents(container, events) {
        if (!events || events.error) {
            container.innerHTML = `<div class="error">Ошибка: ${events?.error || 'Нет данных'}</div>`;
            return;
        }
        if (events.length === 0) {
            container.innerHTML = '<div class="no-events">Событий нет</div>';
            return;
        }
        
        container.innerHTML = events.map(event => `
            <div class="event-card">
                <h3>${event.name}</h3>
                <div class="event-info">
                    <strong>ID:</strong> ${event.id} | 
                    <strong>Дата:</strong> ${new Date(event.date).toLocaleString()} | 
                    <strong>Места:</strong> ${event.freeSeats} / ${event.capacity}
                </div>
                ${event.freeSeats === 0 ? '<div class="sold-out">Мест нет</div>' : '<div class="available">Есть свободные места</div>'}
                <button onclick="selectEvent(${event.id})" class="select-event-btn">
                    Выбрать это мероприятие
                </button>
            </div>
        `).join('');
    },

    renderBookings(container, bookings, eventId = null) {
        console.log('=== RENDER BOOKINGS START ===');
        console.log('Bookings data:', bookings);
        
        container.innerHTML = '';

        if (!bookings || bookings.error) {
            container.innerHTML = `<div class="error">Ошибка: ${bookings?.error || 'Нет данных'}</div>`;
            return;
        }
        if (bookings.length === 0) {
            container.innerHTML = '<div class="no-bookings">Бронирований нет</div>';
            return;
        }

        const now = new Date();
        console.log('Current time:', now);
        
        container.innerHTML = bookings.map(booking => {
            const expiresAt = new Date(booking.expiresAt);
            const isExpired = !booking.paid && expiresAt < now;
            
            console.log('Processing booking:', booking.id, 'paid:', booking.paid, 'expired:', isExpired);
            
            let statusText = '';
            if (booking.paid) {
                statusText = '<span style="color: green; font-weight: bold;">✅ ОПЛАЧЕНО</span>';
            } else if (isExpired) {
                statusText = '<span style="color: red; font-weight: bold;">❌ ПРОСРОЧЕНО</span>';
            } else {
                statusText = `<span style="color: orange; font-weight: bold;">⏳ Ожидает оплаты</span>`;
            }

            // КНОПКА ДОБАВЛЯЕТСЯ ВСЕГДА КОГДА НЕ ОПЛАЧЕНО И НЕ ПРОСРОЧЕНО
            let buttonHtml = '';
            if (!booking.paid && !isExpired) {
                console.log('ADDING BUTTON for booking:', booking.id);
                buttonHtml = `
                    <button class="confirm-btn" onclick="window.confirmBooking(${booking.id}, ${eventId})" 
                            style="background: #28a745; color: white; padding: 10px 20px; border: none; border-radius: 4px; cursor: pointer; margin-top: 10px; font-size: 16px;">
                        ПОДТВЕРДИТЬ ОПЛАТУ
                    </button>
                `;
            }

            return `
                <div class="booking-card" style="border: 2px solid ${booking.paid ? 'green' : isExpired ? 'red' : 'orange'}; padding: 15px; margin: 10px 0; border-radius: 8px;">
                    <div class="booking-info">
                        <strong>Бронь #${booking.id}</strong> | 
                        Мест: ${booking.seats} | 
                        Статус: ${statusText}
                    </div>
                    <div class="booking-dates">
                        Создано: ${new Date(booking.createdAt).toLocaleString()} | 
                        Истекает: ${expiresAt.toLocaleString()}
                    </div>
                    ${buttonHtml}
                </div>
            `;
        }).join('');
        
        console.log('=== RENDER BOOKINGS END ===');
    },

    startAutoRefresh(container, eventId, interval = 30000) {
        return setInterval(async () => {
            console.log('Автообновление...');
            const bookings = await this.fetchBookings(eventId);
            this.renderBookings(container, bookings, eventId);
        }, interval);
    }
};

// ✅ ВАЖНО: Глобальная функция ДОЛЖНА быть определена
window.confirmBooking = async function(bookingId, eventId) {
    console.log('Confirming booking:', bookingId);
    const result = await EventAPI.confirmBooking(bookingId);
    if (result.error) {
        alert('Ошибка: ' + result.error);
    } else {
        alert('Бронь подтверждена! Обновляю список...');
        // Обновляем список
        setTimeout(async () => {
            const container = document.getElementById('my-bookings');
            const newBookings = await EventAPI.fetchBookings(eventId);
            EventAPI.renderBookings(container, newBookings, eventId);
        }, 1000);
    }
};