import { Component, OnInit } from '@angular/core';
import {CommonModule} from '@angular/common';

@Component({
  selector: 'app-refugio',
  standalone:true,
  imports: [ CommonModule],
  templateUrl: './refugio.component.html',
  styleUrl: './refugio.component.css'
})
export class RefugioComponent {
  slides = [
    {img: "/assets/rifugio/3.jpg"},
    {img: "/assets/rifugio/2.jpg"},
    {img: "/assets/rifugio/1.jpg"},
    {img: "/assets/rifugio/2.jpg"}
  ];

}
