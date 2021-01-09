
import { useEffect, useState } from "react";

const useWindowSize = () =>  {
    const [windowSize, setWindowSize] = useState({
      width: 1000,
      height: 1000,
    });
  
    useEffect(() => {
      function handleResize() {
        setWindowSize({
          width: window.innerWidth,
          height: window.innerHeight,
        });
      }
      
      window.addEventListener("resize", handleResize);
      
      handleResize();
      
      return () => window.removeEventListener("resize", handleResize);
    }, []); 
  
    return windowSize;
  }

  export default useWindowSize