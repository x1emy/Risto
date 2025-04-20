import { Component } from '@angular/core';
import { RouterOutlet, RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';
import AOS from 'aos';
import {NavbarComponent} from "./navbar/navbar.component";
import {RestarauntsComponent} from './restaraunts/restaraunts.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet, RouterLink, NavbarComponent, RestarauntsComponent], // 👈 добавь CommonModule
  template: `
    // <nav>
      <a routerLink="/register">Регистрация</a> |
      <a routerLink="/home">Домой</a>
    </nav>-->
    <router-outlet></router-outlet>
    <app-navbar></app-navbar>
    <app-restaraunts></app-restaraunts>
  `,
})
export class AppComponent {

  ngOnInit() {
    AOS.init();
  }

  title = 'booking-frontend';
}
