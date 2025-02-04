import { Component, OnInit } from '@angular/core';
import { ImageService } from "../services/image.service";
import { UserService } from "../services/user.service";

import {
  Chart,
  BarController,
  BarElement,
  CategoryScale,
  LinearScale,
  Title,
  Tooltip,
  Legend,
  PieController,
  LineController,
  RadarController,
  PointElement,
  LineElement,
  RadialLinearScale,
  ArcElement,
  ChartTypeRegistry
} from 'chart.js';
import {Image} from "../objects/image";
import {TopTag} from "../objects/topTag";
import {TopActiveUser} from "../objects/topActiveUser";

@Component({
  selector: 'app-statistics',
  templateUrl: './statistics.component.html',
  styleUrls: ['./statistics.component.scss']
})
export class StatisticsComponent implements OnInit {
  chart: Chart | undefined;
  totalUsers: number = 0;
  totalImages: number = 0;
  totalChartTypes: number = 3;
  selectedChartType: string = 'bar';

  warmColors: string[] = [
    'rgba(173, 216, 230, 0.7)',
    'rgba(144, 238, 144, 0.7)',
    'rgba(255, 182, 193, 0.7)',
    'rgba(255, 239, 213, 0.7)',
    'rgba(221, 160, 221, 0.7)'
  ];

  borderWarmColors: string[] = [
    'rgba(173, 216, 230, 1)',
    'rgba(144, 238, 144, 1)',
    'rgba(255, 182, 193, 1)',
    'rgba(255, 239, 213, 1)',
    'rgba(221, 160, 221, 1)'
  ];

  constructor(private imageService: ImageService,
              private userService: UserService,) {
    Chart.register(
      BarController,
      BarElement,
      CategoryScale,
      LinearScale,
      Title,
      Tooltip,
      Legend,
      PieController,
      LineController,
      RadarController,
      PointElement,
      LineElement,
      RadialLinearScale,
      ArcElement
    );
  }

  ngOnInit(): void {
    this.getAllUsers();
    this.getAllImages();
    this.renderChart('bar');
  }

  getAllUsers(): void {
    this.userService.getUsers().subscribe((users: any[]) => {
      this.totalUsers = users.length;
    });
  }

  getAllImages(): void {
    this.imageService.getImages().subscribe((images: Image[]) => {
      const publicImages = images.filter(image => !image.privateImg);
      this.totalImages = publicImages.length;
    });
  }

  renderChart(type: string): void {
    this.selectedChartType = type;
    const canvas = document.getElementById('dynamicChart') as HTMLCanvasElement;

    if (!canvas) {
      console.error('Canvas element not found!');
      return;
    }

    if (this.chart) {
      this.chart.destroy();
    }

    switch (type) {
      case 'pie':
        this.renderPieChart(canvas);
        break;
      case 'doughnut':
        this.renderDoughnutChart(canvas);
        break;
      case 'bar':
        this.renderBarChart(canvas);
        break;
      default:
        console.error(`Chart type "${type}" is not supported.`);
    }
  }

  private renderPieChart(canvas: HTMLCanvasElement): void {
    this.imageService.getTopUsedTags().subscribe((tags: TopTag[]) => {
      const labels = tags.map(tag => tag.name);
      const data = tags.map(tag => tag.count);
      this.createChart(canvas, 'pie', labels, data, 'Top 5 Most Used Tags');
    });
  }

  private renderDoughnutChart(canvas: HTMLCanvasElement): void {
    this.imageService.getTopActiveUsers().subscribe((users: TopActiveUser[]) => {
      const labels = users.map(user => user.name);
      const data = users.map(user => user.count);
      this.createChart(canvas, 'doughnut', labels, data, 'Top Active Users');
    });
  }

  private renderBarChart(canvas: HTMLCanvasElement): void {
    this.imageService.getTopLikedImages().subscribe((images: Image[]) => {
      const data = images.map(image => image.likes);
      const labels = images.map(image => image.user?.username || 'Untitled');
      this.createChart(canvas, 'bar', labels, data, 'Top Liked Images', true);
    });
  }

  private createChart(
    canvas: HTMLCanvasElement,
    type: keyof ChartTypeRegistry,
    labels: string[],
    data: number[],
    chartLabel: string,
    isBarChart: boolean = false
  ): void {
    const backgroundColors = this.warmColors.slice(0, labels.length);
    const borderColors = this.borderWarmColors.slice(0, labels.length);

    const options: any = {
      responsive: true,
      plugins: {
        legend: {
          position: 'top',
          labels: {
            color: 'black',
            font: {
              size: 14
            }
          }
        },
        tooltip: {
          callbacks: {
            label: (tooltipItem: any) => {
              const itemLabel = labels[tooltipItem.dataIndex];
              const itemCount = data[tooltipItem.dataIndex];
              return `${itemLabel}: ${itemCount}`;
            }
          }
        }
      }
    };

    if (isBarChart) {
      options.scales = {
        y: {
          ticks: {
            stepSize: 1,
            callback: (value: number) => (Number.isInteger(value) ? value : '')
          },
          beginAtZero: true
        }
      };
    }

    this.chart = new Chart(canvas, {
      type: type,
      data: {
        labels: labels,
        datasets: [
          {
            label: chartLabel,
            data: data,
            backgroundColor: backgroundColors,
            borderColor: borderColors,
            borderWidth: 1
          }
        ]
      },
      options: options
    });
  }
}
