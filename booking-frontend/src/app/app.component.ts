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
     <div class="layout">
    <app-navbar></app-navbar>
     <main class="main-content">
        <router-outlet></router-outlet>
     </main>
    </div>

    <app-restaraunts></app-restaraunts>
  `,
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  ngOnInit() {
    AOS.init();
  }

  title = 'booking-frontend';
}
