import CircularProgressBar from '@/app/components/circular_progress_bar.tsx';
import { render, screen } from '@testing-library/react'

describe('Progressbar snapshots', () => {
  it('renders the progressbar.', () => {
    expect(true).toBeTruthy()
    const tree = render(
      <CircularProgressBar
        maxValue={8}
        selectedValue={2}
        radius={180}
        strokeWidth={12}
        label=''
        activeStrokeColor='#05a168'
        inactiveStrokeColor='#ddd'
        backgroundColor='#fff'
        textColor='#ddd'
        labelFontSize={12}
        valueFontSize={60}
        withGradient={false}
        anticlockwise={false}
        initialAngularDisplacement={0}/>)
      expect(tree).toMatchInlineSnapshot(`
{
  "asFragment": [Function],
  "baseElement": <body>
    <div>
      <svg
        height="360"
        width="360"
      >
        <path
          d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
          fill="#05a168"
          style="opacity: 0;"
          transform="rotate(315, 180, 180)"
        />
        <path
          d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
          fill="#05a168"
          style="opacity: 1;"
          transform="rotate(0, 180, 180)"
        />
        <path
          d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
          fill="#05a168"
          style="opacity: 1;"
          transform="rotate(-315, 180, 180)"
        />
        <path
          d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
          fill="#ddd"
          style="opacity: 1;"
          transform="rotate(-630, 180, 180)"
        />
        <path
          d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
          fill="#ddd"
          style="opacity: 1;"
          transform="rotate(-945, 180, 180)"
        />
        <path
          d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
          fill="#ddd"
          style="opacity: 1;"
          transform="rotate(-1260, 180, 180)"
        />
        <path
          d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
          fill="#ddd"
          style="opacity: 1;"
          transform="rotate(-1575, 180, 180)"
        />
        <path
          d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
          fill="#ddd"
          style="opacity: 1;"
          transform="rotate(-1890, 180, 180)"
        />
        <path
          d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
          fill="#ddd"
          style="opacity: 1;"
          transform="rotate(-2205, 180, 180)"
        />
        <circle
          cx="180"
          cy="180"
          fill="#fff"
          r="168"
        />
        <text
          fill="#ddd"
          font-size="60"
          font-weight="bold"
          text-anchor="middle"
          x="180"
          y="200"
        >
          2
        </text>
      </svg>
    </div>
  </body>,
  "container": <div>
    <svg
      height="360"
      width="360"
    >
      <path
        d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
        fill="#05a168"
        style="opacity: 0;"
        transform="rotate(315, 180, 180)"
      />
      <path
        d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
        fill="#05a168"
        style="opacity: 1;"
        transform="rotate(0, 180, 180)"
      />
      <path
        d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
        fill="#05a168"
        style="opacity: 1;"
        transform="rotate(-315, 180, 180)"
      />
      <path
        d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
        fill="#ddd"
        style="opacity: 1;"
        transform="rotate(-630, 180, 180)"
      />
      <path
        d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
        fill="#ddd"
        style="opacity: 1;"
        transform="rotate(-945, 180, 180)"
      />
      <path
        d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
        fill="#ddd"
        style="opacity: 1;"
        transform="rotate(-1260, 180, 180)"
      />
      <path
        d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
        fill="#ddd"
        style="opacity: 1;"
        transform="rotate(-1575, 180, 180)"
      />
      <path
        d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
        fill="#ddd"
        style="opacity: 1;"
        transform="rotate(-1890, 180, 180)"
      />
      <path
        d="M180 180 L180 0 A180 180 0 0 1 294.73631815476415 41.30761630035792 Z"
        fill="#ddd"
        style="opacity: 1;"
        transform="rotate(-2205, 180, 180)"
      />
      <circle
        cx="180"
        cy="180"
        fill="#fff"
        r="168"
      />
      <text
        fill="#ddd"
        font-size="60"
        font-weight="bold"
        text-anchor="middle"
        x="180"
        y="200"
      >
        2
      </text>
    </svg>
  </div>,
  "debug": [Function],
  "findAllByAltText": [Function],
  "findAllByDisplayValue": [Function],
  "findAllByLabelText": [Function],
  "findAllByPlaceholderText": [Function],
  "findAllByRole": [Function],
  "findAllByTestId": [Function],
  "findAllByText": [Function],
  "findAllByTitle": [Function],
  "findByAltText": [Function],
  "findByDisplayValue": [Function],
  "findByLabelText": [Function],
  "findByPlaceholderText": [Function],
  "findByRole": [Function],
  "findByTestId": [Function],
  "findByText": [Function],
  "findByTitle": [Function],
  "getAllByAltText": [Function],
  "getAllByDisplayValue": [Function],
  "getAllByLabelText": [Function],
  "getAllByPlaceholderText": [Function],
  "getAllByRole": [Function],
  "getAllByTestId": [Function],
  "getAllByText": [Function],
  "getAllByTitle": [Function],
  "getByAltText": [Function],
  "getByDisplayValue": [Function],
  "getByLabelText": [Function],
  "getByPlaceholderText": [Function],
  "getByRole": [Function],
  "getByTestId": [Function],
  "getByText": [Function],
  "getByTitle": [Function],
  "queryAllByAltText": [Function],
  "queryAllByDisplayValue": [Function],
  "queryAllByLabelText": [Function],
  "queryAllByPlaceholderText": [Function],
  "queryAllByRole": [Function],
  "queryAllByTestId": [Function],
  "queryAllByText": [Function],
  "queryAllByTitle": [Function],
  "queryByAltText": [Function],
  "queryByDisplayValue": [Function],
  "queryByLabelText": [Function],
  "queryByPlaceholderText": [Function],
  "queryByRole": [Function],
  "queryByTestId": [Function],
  "queryByText": [Function],
  "queryByTitle": [Function],
  "rerender": [Function],
  "unmount": [Function],
}
`);
  })
})
