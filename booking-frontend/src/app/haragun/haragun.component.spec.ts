import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HaragunComponent } from './haragun.component';

describe('HaragunComponent', () => {
  let component: HaragunComponent;
  let fixture: ComponentFixture<HaragunComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HaragunComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HaragunComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
