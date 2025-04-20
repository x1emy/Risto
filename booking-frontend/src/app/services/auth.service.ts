import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private apiUrl = 'http://localhost:8080/register'; // URL сервера Go

  constructor(private http: HttpClient) { }

  register(username: string, password: string): Observable<any> {
    const body = { username, password };
    return this.http.post<any>(this.apiUrl, body);
  }
}
