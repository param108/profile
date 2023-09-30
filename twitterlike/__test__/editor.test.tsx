import Editor from '@/app/components/editor';
import { render, screen } from '@testing-library/react'

describe('Editor snapshot.', () => {
  it('renders the logged in.', () => {
    expect(true).toBeTruthy()
    const tree = render(<Editor
                        url="https://ui.tribist.com"
                        isLoggedIn={false}
                        defaultMessage={`
This is a blog. A **blog** of _tweets_.
Used to be called **micro-blogging** until twitter
**Hijacked** the space.
`}
                        onSendClicked={()=>{}}
                        onChange={()=>{}}
                        showLoading={false}
                        hideClicked={()=>{}}
                        hideable={false}
                        value=""
/>)
    expect(tree).toMatchInlineSnapshot(`
{
  "asFragment": [Function],
  "baseElement": <body>
    <div>
      <div
        class="bg-sky-200 w-[96%] md:w-[510px] rounded mb-[10px]"
      >
        <div
          class="w-[96%] md:w-[500px] h-[150px] bg-white mt-[5px] mx-[2%] md:mx-[5px] p-[5px] rounded focus:outline-none overflow-x-auto text-gray-600"
        >
          <div>
            <div
              class=""
            >
              <p>
                This is a blog. A 
                <strong>
                  blog
                </strong>
                 of 
                <em>
                  tweets
                </em>
                .
Used to be called 
                <strong>
                  micro-blogging
                </strong>
                 until twitter

                <strong>
                  Hijacked
                </strong>
                 the space.
              </p>
            </div>
          </div>
        </div>
        <div
          class="block pt-[5px] min-h-[35px]"
        />
      </div>
    </div>
  </body>,
  "container": <div>
    <div
      class="bg-sky-200 w-[96%] md:w-[510px] rounded mb-[10px]"
    >
      <div
        class="w-[96%] md:w-[500px] h-[150px] bg-white mt-[5px] mx-[2%] md:mx-[5px] p-[5px] rounded focus:outline-none overflow-x-auto text-gray-600"
      >
        <div>
          <div
            class=""
          >
            <p>
              This is a blog. A 
              <strong>
                blog
              </strong>
               of 
              <em>
                tweets
              </em>
              .
Used to be called 
              <strong>
                micro-blogging
              </strong>
               until twitter

              <strong>
                Hijacked
              </strong>
               the space.
            </p>
          </div>
        </div>
      </div>
      <div
        class="block pt-[5px] min-h-[35px]"
      />
    </div>
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

  it('renders the logged out', () => {
    expect(true).toBeTruthy()
    const tree = render(<Editor
                          url="https://ui.tribist.com"
                          isLoggedIn={true}
                          defaultMessage={`
This is a blog. A **blog** of _tweets_.
Used to be called **micro-blogging** until twitter
**Hijacked** the space.
`}
                        onSendClicked={()=>{}}
                        onChange={()=>{}}
                        showLoading={false}
                        hideClicked={()=>{}}
                        hideable={false}
                        value=""
      />)
    expect(tree).toMatchInlineSnapshot(`
{
  "asFragment": [Function],
  "baseElement": <body>
    <div>
      <div
        class="bg-sky-200 w-[96%] md:w-[510px] rounded mb-[10px]"
      >
        <div>
          <textarea
            class="block dark:bg-slate-700 w-[96%] md:w-[500px] h-[150px] resize-none caret-red-500 mt-[5px] mx-[2%] md:mx-[5px] pl-[10px] pr-[5px] py-[5px] rounded focus:outline-none text-black dark:text-slate-100"
            placeholder="What are you thinking about ?"
          />
          <textarea
            class="block dark:bg-slate-700 w-[96%] md:w-[500px] h-[50px] resize-none caret-red-500 mt-[5px] mx-[2%] md:mx-[5px] pl-[10px] pr-[5px] py-[5px] rounded focus:outline-none text-black dark:text-slate-100"
            placeholder="display tags..."
          />
        </div>
        <div
          class="block pt-[5px]"
        >
          <button
            class="px-[10px]"
          >
            <svg
              class="text-indigo-800"
              fill="none"
              height="30"
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              viewBox="0 0 24 24"
              width="30"
              xmlns="http://www.w3.org/2000/svg"
            >
              <rect
                height="18"
                rx="2"
                ry="2"
                width="18"
                x="3"
                y="3"
              />
              <circle
                cx="8.5"
                cy="8.5"
                r="1.5"
              />
              <polyline
                points="21 15 16 10 5 21"
              />
            </svg>
          </button>
          <div
            class="inline-block float-right pr-[10px]"
          />
          <div
            class="inline-block float-right"
          >
            <button
              class="px-[10px] float-right"
            >
              <svg
                class="text-indigo-800"
                fill="none"
                height="30"
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                viewBox="0 0 24 24"
                width="30"
                xmlns="http://www.w3.org/2000/svg"
              >
                <line
                  x1="22"
                  x2="11"
                  y1="2"
                  y2="13"
                />
                <polygon
                  points="22 2 15 22 11 13 2 9 22 2"
                />
              </svg>
            </button>
            <span
              class="text-indigo-800 text-lg float-right items-center px-[10px]"
            >
              50
            </span>
          </div>
        </div>
      </div>
    </div>
  </body>,
  "container": <div>
    <div
      class="bg-sky-200 w-[96%] md:w-[510px] rounded mb-[10px]"
    >
      <div>
        <textarea
          class="block dark:bg-slate-700 w-[96%] md:w-[500px] h-[150px] resize-none caret-red-500 mt-[5px] mx-[2%] md:mx-[5px] pl-[10px] pr-[5px] py-[5px] rounded focus:outline-none text-black dark:text-slate-100"
          placeholder="What are you thinking about ?"
        />
        <textarea
          class="block dark:bg-slate-700 w-[96%] md:w-[500px] h-[50px] resize-none caret-red-500 mt-[5px] mx-[2%] md:mx-[5px] pl-[10px] pr-[5px] py-[5px] rounded focus:outline-none text-black dark:text-slate-100"
          placeholder="display tags..."
        />
      </div>
      <div
        class="block pt-[5px]"
      >
        <button
          class="px-[10px]"
        >
          <svg
            class="text-indigo-800"
            fill="none"
            height="30"
            stroke="currentColor"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            viewBox="0 0 24 24"
            width="30"
            xmlns="http://www.w3.org/2000/svg"
          >
            <rect
              height="18"
              rx="2"
              ry="2"
              width="18"
              x="3"
              y="3"
            />
            <circle
              cx="8.5"
              cy="8.5"
              r="1.5"
            />
            <polyline
              points="21 15 16 10 5 21"
            />
          </svg>
        </button>
        <div
          class="inline-block float-right pr-[10px]"
        />
        <div
          class="inline-block float-right"
        >
          <button
            class="px-[10px] float-right"
          >
            <svg
              class="text-indigo-800"
              fill="none"
              height="30"
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              viewBox="0 0 24 24"
              width="30"
              xmlns="http://www.w3.org/2000/svg"
            >
              <line
                x1="22"
                x2="11"
                y1="2"
                y2="13"
              />
              <polygon
                points="22 2 15 22 11 13 2 9 22 2"
              />
            </svg>
          </button>
          <span
            class="text-indigo-800 text-lg float-right items-center px-[10px]"
          >
            50
          </span>
        </div>
      </div>
    </div>
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
