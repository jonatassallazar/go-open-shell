import { SetStateAction, useEffect, useRef, useState } from 'react';
import { IconSend } from '@tabler/icons-react';
import {
  ActionIcon,
  Code,
  Container,
  Flex,
  ScrollArea,
  Tabs,
  Text,
  TextInput,
  Timeline,
} from '@mantine/core';
import { useDebouncedCallback } from '@mantine/hooks';
import { CreateSession, SendMessageInput } from '../../wailsjs/go/main/App';
import { terminal } from '../../wailsjs/go/models';
import { EventsEmit, EventsOn } from '../../wailsjs/runtime';

const Output = (text: string) => {
  const lines: string[] = text.split('\n');

  return lines.map((line, index) => <Text key={`text_${index}`}>{line}</Text>);
};

function App() {
  const [session, setSession] = useState<terminal.TerminalSession>();
  const [disabled, setDisabled] = useState<boolean>(true);
  const [input, setInput] = useState<string>('');

  EventsOn('session-initiated', () => {
    console.log('session-initiated');
    setDisabled(false);
  });

  useEffect(() => {
    CreateSession()
      .then((newSession) => {
        setSession(newSession);
        SendMessageInput(newSession!.id)
          .then()
          .catch((err) => {
            console.error(err);
          });
      })
      .catch((err) => {
        console.error(err);
      });
  }, []);

  EventsOn('input-response', (newSession) => {
    setSession(newSession);
    setInput('');

    handleScrollDebounced();
  });

  const sendInput = async () => {
    // const inputEvent: terminal.InputEvent = new terminal.InputEvent({
    //   in: input,
    //   out: '',
    //   dir: '',
    //   time: null,
    // });

    EventsEmit('send-input', input);
    console.log('send-input');
  };

  const handleKeyPress = (e: any) => {
    if (e.key === 'Enter') {
      sendInput();
    }
  };

  const viewport = useRef<HTMLDivElement>(null);

  const handleScrollDebounced = useDebouncedCallback(() => {
    // const scrollToBottom = () => {
    if (!viewport) {
      return;
    }
    viewport.current!.scrollTo({ top: viewport.current!.scrollHeight, behavior: 'smooth' });
    // };
  }, 100);

  return (
    <div id="App">
      <Tabs>
        <Tabs.List>
          <Tabs.Tab value="Console1">Console 1</Tabs.Tab>
          <Tabs.Tab value="Console2">Console 2</Tabs.Tab>
          <Tabs.Tab value="Console3">Console 3</Tabs.Tab>
        </Tabs.List>
      </Tabs>
      <Container my="md">
        <ScrollArea h="500px" type="auto" offsetScrollbars scrollbarSize={6} viewportRef={viewport}>
          {/* <Timeline>
            {session?.history?.map((line, index) => (
              <Timeline.Item key={`line_${index}`} title={line.time}>
                {Output(line.out)}
              </Timeline.Item>
            ))}
          </Timeline> */}
          {Output(session?.input.out || '')}
          {/* <Code>{session?.input.out}</Code> */}
        </ScrollArea>
        <Flex mih={100} gap="md" direction="row" justify="space-between" align="end">
          <TextInput
            style={{ width: '100%' }}
            label="Shell Input"
            description={`Directory: ${session?.dir}`}
            value={input}
            onChange={(e: { currentTarget: { value: SetStateAction<string> } }) =>
              setInput(e.currentTarget.value)
            }
            onKeyDown={(e) => handleKeyPress(e)}
            disabled={disabled}
          />
          <ActionIcon
            variant="filled"
            aria-label="Send"
            size="lg"
            onClick={sendInput}
            disabled={disabled}
          >
            <IconSend />
          </ActionIcon>
        </Flex>
      </Container>
    </div>
  );
}

export default App;
