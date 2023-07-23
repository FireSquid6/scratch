import Component from "./Component";
import { render } from "@testing-library/react";

describe("Component", () => {
  it("renders without crashing", () => {
    render(<Component />);
  });
});
