import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [],
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class HomeComponent implements OnInit {
  currentDate: Date = new Date();
  currentMonthName: string = '';
  currentYear: number = 0;
  weekDays: string[] = ['BC', 'ПН', 'ВТ', 'СР', 'ЧТ', 'ПТ', 'СБ'];
  calendarGrid: number[][] = [];

  ngOnInit(): void {
    this.updateCalendar();
  }

  updateCalendar(): void {
    const monthNames = ["Январь", "Февраль", "Март", "Апрель", "Май", "Июнь",
      "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"];

    this.currentMonthName = monthNames[this.currentDate.getMonth()];
    this.currentYear = this.currentDate.getFullYear();
    this.generateCalendarGrid();
  }

  generateCalendarGrid(): void {
    // Здесь должна быть логика генерации календарной сетки
    // Это упрощенный пример - в реальном проекте нужно правильно вычислять дни
    this.calendarGrid = [
      [27, 28, 29, 30, 31, 1, 2],
      [3, 4, 5, 6, 7, 8, 9],
      [10, 11, 12, 13, 14, 15, 16],
      [17, 18, 19, 20, 21, 22, 23],
      [24, 25, 26, 27, 28, 29, 30]
    ];
  }

  previousMonth(): void {
    this.currentDate = new Date(
      this.currentDate.getFullYear(),
      this.currentDate.getMonth() - 1,
      1
    );
    this.updateCalendar();
  }

  nextMonth(): void {
    this.currentDate = new Date(
      this.currentDate.getFullYear(),
      this.currentDate.getMonth() + 1,
      1
    );
    this.updateCalendar();
  }
}
