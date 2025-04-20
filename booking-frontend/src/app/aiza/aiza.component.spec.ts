import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AizaComponent } from './aiza.component';

describe('AizaComponent', () => {
  let component: AizaComponent;
  let fixture: ComponentFixture<AizaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AizaComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AizaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
