import { Component, ViewEncapsulation, ElementRef, ViewChild, AfterViewInit } from '@angular/core';
import { RouterLink } from '@angular/router'; // 👈 ВАЖНО: импортируем RouterLink

@Component({
  selector: 'app-navbar',
  standalone: true,
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css'],
  encapsulation: ViewEncapsulation.None,
  imports: [RouterLink]
})
export class NavbarComponent implements AfterViewInit {
  @ViewChild('toggleBtn', { static: false }) toggleButtonRef!: ElementRef;
  @ViewChild('sidebar', { static: false }) sidebarRef!: ElementRef;


  ngAfterViewInit(): void {
    // Здесь можно инициализировать поведение, если нужно
  }

  toggleSidebar(): void {
    const sidebar = this.sidebarRef.nativeElement as HTMLElement;
    const toggleButton = this.toggleButtonRef.nativeElement as HTMLElement;

    sidebar.classList.toggle('close');
    toggleButton.classList.toggle('rotate');

    this.closeAllSubMenus();
  }

  toggleSubMenu(button: HTMLElement): void {
    const sidebar = this.sidebarRef.nativeElement as HTMLElement;
    const toggleButton = this.toggleButtonRef.nativeElement as HTMLElement;
    const subMenu = button.nextElementSibling as HTMLElement;

    if (!subMenu.classList.contains('show')) {
      this.closeAllSubMenus();
    }

    subMenu.classList.toggle('show');
    button.classList.toggle('rotate');

    if (sidebar.classList.contains('close')) {
      sidebar.classList.remove('close');
      toggleButton.classList.remove('rotate');
    }
  }

  closeAllSubMenus(): void {
    const sidebar = this.sidebarRef.nativeElement as HTMLElement;
    const shownMenus = sidebar.querySelectorAll('.show');

    shownMenus.forEach(ul => {
      ul.classList.remove('show');
      const button = ul.previousElementSibling as HTMLElement;
      if (button) {
        button.classList.remove('rotate');
      }
    });
  }
}
