import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BestiaComponent } from './bestia.component';

describe('BestiaComponent', () => {
  let component: BestiaComponent;
  let fixture: ComponentFixture<BestiaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BestiaComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BestiaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
