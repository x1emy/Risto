import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { HttpClientModule, HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [CommonModule, FormsModule, HttpClientModule],
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  username = '';
  password = '';
  responseMessage = '';

  constructor(private http: HttpClient) {}

  register() {
    const body = { username: this.username, password: this.password };
    this.http.post('http://localhost:8080/register', body).subscribe({
      next: () => this.responseMessage = 'Успешная регистрация!',
      error: () => this.responseMessage = 'Ошибка при регистрации.'
    });
  }
}
