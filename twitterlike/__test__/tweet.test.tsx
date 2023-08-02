import Tweet from '@/app/components/tweet';
import { render, screen } from '@testing-library/react'

describe('Tweet snapshots', () => {
  it('renders the correct message.', () => {
    expect(true).toBeTruthy()
    const tree = render(<Tweet
                          tweet_id={"1"}
                          tweet={`
You can do all this here and you own your data,
download as you wish.

Unlike twitter this is not a **performance**,

this is **recreation**. This is **expression**.

This is **Freedom**!

_Interested ?_

Then [**signup**](https://google.com)!
`}
                        key={1}
                        date={"13/03/2022"}
                        onClick={()=>{}}
                        deleteClicked={()=>{}}
                        editClicked={()=>{}}
                        showMenu={false}
                        url="https://ui.tribist.com"
                        />)
      expect(tree).toMatchInlineSnapshot(`
{
  "asFragment": [Function],
  "baseElement": <body>
    <div>
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
            >
              13/03/2022
            </i>
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
                  You can do all this here and you own your data,
download as you wish.
                </p>
                

                <p>
                  Unlike twitter this is not a 
                  <strong>
                    performance
                  </strong>
                  ,
                </p>
                

                <p>
                  this is 
                  <strong>
                    recreation
                  </strong>
                  . This is 
                  <strong>
                    expression
                  </strong>
                  .
                </p>
                

                <p>
                  This is 
                  <strong>
                    Freedom
                  </strong>
                  !
                </p>
                

                <p>
                  <em>
                    Interested ?
                  </em>
                </p>
                

                <p>
                  Then 
                  <a
                    href="https://google.com"
                  >
                    <strong>
                      signup
                    </strong>
                  </a>
                  !
                </p>
              </div>
            </div>
          </span>
        </div>
      </div>
    </div>
  </body>,
  "container": <div>
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
          >
            13/03/2022
          </i>
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
                You can do all this here and you own your data,
download as you wish.
              </p>
              

              <p>
                Unlike twitter this is not a 
                <strong>
                  performance
                </strong>
                ,
              </p>
              

              <p>
                this is 
                <strong>
                  recreation
                </strong>
                . This is 
                <strong>
                  expression
                </strong>
                .
              </p>
              

              <p>
                This is 
                <strong>
                  Freedom
                </strong>
                !
              </p>
              

              <p>
                <em>
                  Interested ?
                </em>
              </p>
              

              <p>
                Then 
                <a
                  href="https://google.com"
                >
                  <strong>
                    signup
                  </strong>
                </a>
                !
              </p>
            </div>
          </div>
        </span>
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
