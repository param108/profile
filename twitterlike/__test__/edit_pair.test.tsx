import EditPair from '@/app/components/edit_tweet_pair';
import { render, screen } from '@testing-library/react'

describe('Editpair snapshot.', () => {
  it('renders the editting and view view', () => {
    expect(true).toBeTruthy()
    const tree = render(<EditPair
			                  editting={true}
			                  viewing={true}
                        isLoggedIn={true}
			                  tweet={null}
                        onChange={()=>{}}
			                  onSendClicked={()=>{}}
			                  value={"Hello World"}
			                  showLoading={true}
                        defaultMessage={`
This is a blog. A **blog** of _tweets_.
Used to be called **micro-blogging** until twitter
**Hijacked** the space.
`}
                        editClicked={()=>{}}
                        deleteClicked={()=>{}}
                        hideClicked={()=>{}}
                        editorHideable={false}
                        url="https://ui.tribist.com"
                        showMenu={false}
      />)
    expect(tree).toMatchInlineSnapshot(`
{
  "asFragment": [Function],
  "baseElement": <body>
    <div>
      <div
        class="w-full flex flex-col items-center invisible"
      >
        <div
          class="bg-sky-200 w-[96%] md:w-[510px] rounded mb-[10px]"
        >
          <div>
            <textarea
              class="block dark:bg-slate-700 w-[96%] md:w-[500px] h-[150px] resize-none caret-red-500 mt-[5px] mx-[2%] md:mx-[5px] pl-[10px] pr-[5px] py-[5px] rounded focus:outline-none text-black dark:text-slate-100"
              placeholder="What are you thinking about ?"
            >
              Hello World
            </textarea>
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
            >
              <span
                class="inline-block"
                style="display: inherit; width: 30px; height: 30px; position: relative;"
              >
                <span
                  style="position: absolute; top: 0px; left: 0px; width: 30px; height: 30px; border: 3px solid #ec4899; opacity: 0.4; border-radius: 100%; animation-fill-mode: forwards; animation: react-spinners-RingLoader-right 2s 0s infinite linear;"
                />
                <span
                  style="position: absolute; top: 0px; left: 0px; width: 30px; height: 30px; border: 3px solid #ec4899; opacity: 0.4; border-radius: 100%; animation-fill-mode: forwards; animation: react-spinners-RingLoader-left 2s 0s infinite linear;"
                />
              </span>
            </div>
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
                489
              </span>
            </div>
          </div>
        </div>
        <div
          class="flex flex-col items-center w-full"
        >
          <div
            class="border w-[90%] md:w-[510px]"
          >
            <div
              class="
            w-full bg-white dark:bg-slate-700 hover:bg-cyan-50
            min-h-[100px] pl-[15px] pr-[5px] pt-[5px]
            pb-[40px] overflow-x-auto m-auto invisible"
            >
              <div
                class="flex flex-row"
              >
                <i
                  class="text-gray-500"
                />
                <br />
              </div>
              <span
                class="text-gray-600 dark:text-slate-100"
              >
                <div>
                  <div
                    class=""
                  >
                    <p>
                      Hello World
                    </p>
                  </div>
                </div>
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </body>,
  "container": <div>
    <div
      class="w-full flex flex-col items-center invisible"
    >
      <div
        class="bg-sky-200 w-[96%] md:w-[510px] rounded mb-[10px]"
      >
        <div>
          <textarea
            class="block dark:bg-slate-700 w-[96%] md:w-[500px] h-[150px] resize-none caret-red-500 mt-[5px] mx-[2%] md:mx-[5px] pl-[10px] pr-[5px] py-[5px] rounded focus:outline-none text-black dark:text-slate-100"
            placeholder="What are you thinking about ?"
          >
            Hello World
          </textarea>
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
          >
            <span
              class="inline-block"
              style="display: inherit; width: 30px; height: 30px; position: relative;"
            >
              <span
                style="position: absolute; top: 0px; left: 0px; width: 30px; height: 30px; border: 3px solid #ec4899; opacity: 0.4; border-radius: 100%; animation-fill-mode: forwards; animation: react-spinners-RingLoader-right 2s 0s infinite linear;"
              />
              <span
                style="position: absolute; top: 0px; left: 0px; width: 30px; height: 30px; border: 3px solid #ec4899; opacity: 0.4; border-radius: 100%; animation-fill-mode: forwards; animation: react-spinners-RingLoader-left 2s 0s infinite linear;"
              />
            </span>
          </div>
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
              489
            </span>
          </div>
        </div>
      </div>
      <div
        class="flex flex-col items-center w-full"
      >
        <div
          class="border w-[90%] md:w-[510px]"
        >
          <div
            class="
            w-full bg-white dark:bg-slate-700 hover:bg-cyan-50
            min-h-[100px] pl-[15px] pr-[5px] pt-[5px]
            pb-[40px] overflow-x-auto m-auto invisible"
          >
            <div
              class="flex flex-row"
            >
              <i
                class="text-gray-500"
              />
              <br />
            </div>
            <span
              class="text-gray-600 dark:text-slate-100"
            >
              <div>
                <div
                  class=""
                >
                  <p>
                    Hello World
                  </p>
                </div>
              </div>
            </span>
          </div>
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

