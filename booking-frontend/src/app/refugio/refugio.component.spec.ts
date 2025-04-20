import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RefugioComponent } from './refugio.component';

describe('RefugioComponent', () => {
  let component: RefugioComponent;
  let fixture: ComponentFixture<RefugioComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RefugioComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RefugioComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
