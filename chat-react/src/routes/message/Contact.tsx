import { useEffect } from "react";
import Header from "../../components/contact/Header";
import { useMainContext } from "../../contexts/MainContext";
import ContactGroup from "../../components/contact/contact-group/ContactGroup";

export default function Contact() {
    const { setType } = useMainContext();

    useEffect(() => {
        setType("conversation");
    }, []);

    return (
        <div className="h-full w-full lg:w-72 border-r border-r-neutral-200 dark:border-r-neutral-700 flex flex-col lg:flex-shrink-0 bg-white dark:bg-neutral-800 peer/conversation no-conversation">
            <Header />
            <ContactGroup />
        </div>
    );
}
